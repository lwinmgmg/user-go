package controller

import (
	"errors"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
)

var (
	ErrEmailAldyConfirmed = errors.New("email_already_confirmed")
)

func (ctrl *Controller) EmailConfirm(userCode string) (loginTkn LoginToken, err error) {
	loginTkn.SendOtpType = SOtpEmail
	loginTkn.TokenType = OTP_TKN
	user := models.User{}
	if _, err = user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	loginTkn.UserCode = user.Code
	if user.Partner.IsEmailConfirmed {
		err = ErrEmailAldyConfirmed
		return
	}
	if err = user.Partner.CheckEmail(ctrl.RoDb); err != nil {
		return
	}
	// Otp Authentication is required
	passCode, err := GenerateOtp(&loginTkn, user.OtpUrl, user.Code, services.OtpEmail, ctrl.RedisCtrl, ctrl.Otp,
		time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, nil)
	if err != nil {
		return
	}
	err = ctrl.LoginMail.Send(passCode, []string{user.Partner.Email})
	return
}
