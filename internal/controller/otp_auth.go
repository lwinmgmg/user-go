package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	redisctrl "github.com/lwinmgmg/user-go/pkg/redis-ctrl"
)

type OtpAuth struct {
	AccessToken string `form:"access_token" binding:"required,min=3"`
	PassCode    string `form:"passcode" binding:"required,min=3"`
}

var (
	ErrUnknownOtpConfirmType = errors.New("unknown_otp_confirm_type")
	ErrFailedToValidateOtp   = errors.New("otp_failed")
)

func GenerateOtp(loginTkn *LoginToken, otpUrl, userCode string, otpConfirmType services.OtpConfirmType,
	redisCtrl *redisctrl.RedisCtrl, otpSer *services.OtpService, duration time.Duration, value map[string]any) (string, error) {
	uuid4 := hashing.NewUuid4() + hashing.NewUuid4() + userCode
	loginTkn.TokenType = OTP_TKN
	loginTkn.AccessToken = string(uuid4)
	otpVal, err := services.EncodeOtpValue(otpUrl, userCode, otpConfirmType, value)
	if err != nil {
		return "", err
	}
	if err := redisCtrl.SetKey(fmt.Sprintf("otp:%v", uuid4), otpVal, duration); err != nil {
		return "", err
	}
	passCode, err := otpSer.GenerateCode(otpUrl)
	if err != nil {
		return "", err
	}
	return passCode, nil
}

func (ctrl Controller) OtpAuth(otpAuth *OtpAuth, user *models.User) (loginTkn LoginToken, otpConfirmType services.OtpConfirmType, err error) {
	loginTkn.TokenType = BEARER
	loginTkn.SendOtpType = SOtpAuth
	otpConfirmType = services.OtpLogin
	val, err := ctrl.RedisCtrl.GetKey(fmt.Sprintf("otp:%v", otpAuth.AccessToken))
	if err != nil {
		return
	}
	defer func() {
		if err == nil {
			err = ctrl.RedisCtrl.DelKey([]string{fmt.Sprintf("otp:%v", otpAuth.AccessToken)})
		}
	}()
	otpValue, err := services.ParseOtpValue(val)
	if err != nil {
		return
	}
	otpConfirmType = otpValue.Type
	if _, err = user.GetPartnerByCode(otpValue.Code, ctrl.RoDb); err != nil {
		return
	}
	loginTkn.UserCode = user.Code
	if !ctrl.Otp.Validate(otpAuth.PassCode, otpValue.Url) {
		err = ErrFailedToValidateOtp
		return
	}
	isUser := true
	switch otpConfirmType {
	case services.OtpLogin:
		formattedKey := services.FormatJwtKey(user.Username, user.Code, string(user.Password), ctrl.Setting.JwtService.Key)
		loginTkn.AccessToken, err = services.GenerateUserLoginJwt(user.Code, formattedKey, ctrl.Setting, ctrl.JwtCtrl)
		return
	case services.OtpEmail:
		user.Partner.IsEmailConfirmed = true
		isUser = false
	case services.OtpAuthr:
		user.IsAuthenticator = true
	case services.OtpEnable:
		user.OtpUrl = otpValue.Url
	case services.OtpPhone:
		user.Partner.IsPhoneConfirmed = true
		isUser = false
	case services.OtpChangePass:
		password, ok := otpValue.Value["password"]
		if !ok {
			err = ErrFailedToValidateOtp
		}
		passwordStr, ok := password.(string)
		if !ok {
			err = ErrFailedToValidateOtp
		}
		var hashPass []byte
		hashPass, err = hashing.Hash256(passwordStr)
		if err != nil {
			return
		}
		user.Password = hashPass
	}
	if isUser {
		err = ctrl.Db.Save(user).Error
	} else {
		err = ctrl.Db.Save(user.Partner).Error
	}

	return
}
