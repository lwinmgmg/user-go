package controller

import (
	"errors"
	"fmt"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
)

type OtpAuth struct {
	AccessToken string `json:"access_token"`
	PassCode    string `json:"passcode"`
}

var (
	ErrUnknownOtpConfirmType = errors.New("unknown_otp_confirm_type")
	ErrFailedToValidateOtp   = errors.New("otp_failed")
)

func (ctrl Controller) OtpAuth(otpAuth *OtpAuth, user *models.User) (otpConfirmType services.OtpConfirmType, err error) {
	otpConfirmType = services.OtpLogin
	val, err := ctrl.RedisCtrl.GetKey(fmt.Sprintf("otp:%v", otpAuth.AccessToken))
	if err != nil {
		return
	}
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
	switch otpConfirmType {
	case services.OtpEmail:
		user.Partner.IsEmailConfirmed = true
	case services.OtpAuthr:
		user.IsAuthenticator = true
	case services.OtpEnable:
		user.OtpUrl = otpValue.Url
	case services.OtpPhone:
		user.Partner.IsPhoneConfirmed = true
	}
	err = ctrl.Db.Save(user).Error
	return
}
