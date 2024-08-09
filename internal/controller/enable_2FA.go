package controller

import (
	"errors"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	otpctrl "github.com/lwinmgmg/user-go/pkg/otp-ctrl"
)

var (
	ErrEmailConfirmRequired = errors.New("email_confirm_required")
	ErrAldyEnable2FA        = errors.New("already_enabled_2fa")
)

func (ctrl *Controller) Enable2FA(userCode string) (loginTkn LoginToken, err error) {
	loginTkn.TokenType = OTP_TKN
	user := models.User{}
	if _, err = user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	loginTkn.UserCode = user.Code
	if !user.Partner.IsEmailConfirmed {
		err = ErrEmailConfirmRequired
	}
	if err == ErrEmailConfirmRequired && !user.Partner.IsPhoneConfirmed {
		return
	}
	if user.OtpUrl != "" {
		err = ErrAldyEnable2FA
		return
	}
	// generate otp url
	url, err := ctrl.Otp.OtpCtrl.GenerateOtpUrl(user.Partner.Email, otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		return
	}
	// Generating Otp
	passCode, err := GenerateOtp(&loginTkn, url, user.Code, services.OtpEnable, ctrl.RedisCtrl, ctrl.Otp,
		time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, nil)
	if err != nil {
		return
	}
	// Send Otp through email or phone
	if user.Partner.IsEmailConfirmed {
		loginTkn.SendOtpType = SOtpEmail
		err = ctrl.LoginMail.Send(passCode, []string{user.Partner.Email})
		if err != nil {
			return
		}
	} else if user.Partner.IsPhoneConfirmed {
		loginTkn.SendOtpType = SOtpPhone
		err = ctrl.PhoneService.Send(passCode, user.Partner.Phone)
		if err != nil {
			return
		}
	}
	return
}
