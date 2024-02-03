package controller

import (
	"errors"
	"fmt"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
)

type OtpAuth struct {
	AccessToken string `json:"access_token" binding:"required,min=3"`
	PassCode    string `json:"passcode" binding:"required,min=3"`
}

var (
	ErrUnknownOtpConfirmType = errors.New("unknown_otp_confirm_type")
	ErrFailedToValidateOtp   = errors.New("otp_failed")
)

func (ctrl Controller) OtpAuth(otpAuth *OtpAuth, user *models.User) (loginTkn LoginToken, otpConfirmType services.OtpConfirmType, err error) {
	loginTkn.TokenType = BEARER
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
	}
	if isUser {
		err = ctrl.Db.Save(user).Error
	} else {
		err = ctrl.Db.Save(user.Partner).Error
	}

	return
}
