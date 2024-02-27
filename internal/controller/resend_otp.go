package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/redis/go-redis/v9"
)

type SendOtpType string

const (
	SOtpPhone  SendOtpType = "phone"
	SOtpEmail  SendOtpType = "email"
	SOtpAuth   SendOtpType = "auth"
	SOtpChPass SendOtpType = "ch_pass"
)

var (
	ErrOtpNotFound          error = errors.New("otp_not_found")
	ErrROtpTypeIsNotConfirm error = errors.New("rotp_type_not_confirm")
)

type ResendOtpRequest struct {
	AccessToken string      `form:"access_token" binding:"required"`
	SendOtpType SendOtpType `form:"sotp_type" binding:"required"`
}

type ResendOtpResponse struct {
	AccessToken string      `json:"access_token"`
	SendOtpType SendOtpType `json:"sotp_type"`
	Message     string      `json:"mesg"`
}

func (ctrl *Controller) ResendOtp(data *ResendOtpRequest) (rsendOtp ResendOtpResponse, err error) {
	rsendOtp.AccessToken = data.AccessToken
	rsendOtp.SendOtpType = data.SendOtpType
	val, err := ctrl.RedisCtrl.GetKey(fmt.Sprintf("otp:%v", data.AccessToken))
	if err != nil {
		if err == redis.Nil {
			err = ErrOtpNotFound
			return
		}
		return
	}
	defer func() {
		if err == nil {
			err = ctrl.RedisCtrl.SetKey(fmt.Sprintf("otp:%v", data.AccessToken), val, 5*time.Minute)
		}
	}()
	otpValue, err := services.ParseOtpValue(val)
	if err != nil {
		return
	}
	passCode, err := ctrl.Otp.GenerateCode(otpValue.Url)
	if err != nil {
		return
	}
	user := models.User{}
	if _, err = user.GetPartnerByCode(otpValue.Code, ctrl.RoDb); err != nil {
		return
	}
	if data.SendOtpType == SOtpEmail && user.Partner.IsEmailConfirmed {
		err = ctrl.LoginMail.Send(passCode, []string{user.Partner.Email})
		rsendOtp.Message = fmt.Sprintf("Otp send to the email [%v...]", user.Partner.Email[:4])
		return
	} else if data.SendOtpType == SOtpPhone && user.Partner.IsPhoneConfirmed {
		err = ctrl.PhoneService.Send(passCode, user.Partner.Phone)
		rsendOtp.Message = fmt.Sprintf("Otp send to the phone [%v...]", user.Partner.Phone[:4])
		return
	}
	err = ErrROtpTypeIsNotConfirm
	return
}
