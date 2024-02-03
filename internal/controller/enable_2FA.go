package controller

import (
	"errors"
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
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
	if !user.Partner.IsEmailConfirmed {
		err = ErrEmailConfirmRequired
		return
	}
	uuid4 := hashing.NewUuid4()
	loginTkn.AccessToken = string(uuid4)
	tknExpTime := 5 * time.Minute
	if user.OtpUrl != "" {
		err = ErrAldyEnable2FA
		return
	}
	url, err := ctrl.Otp.OtpCtrl.GenerateOtpUrl(user.Partner.Email, otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		return
	}
	otpVal, err := services.EncodeOtpValue(url, user.Code, services.OtpEnable)
	if err != nil {
		return
	}
	if err = ctrl.RedisCtrl.SetKey(fmt.Sprintf("otp:%v", uuid4), otpVal, tknExpTime); err != nil {
		return
	}
	passCode, err := ctrl.Otp.GenerateCode(url)
	if err != nil {
		return
	}
	err = ctrl.LoginMail.Send(passCode, []string{user.Partner.Email})
	if err != nil {
		return
	}
	return
}
