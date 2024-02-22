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
	uuid4 := hashing.NewUuid4()
	loginTkn.AccessToken = string(uuid4)
	tknExpTime := time.Duration(ctrl.Setting.OtpService.OtpDuration) * time.Second
	url, err := ctrl.Otp.OtpCtrl.GenerateOtpUrl(user.Partner.Email, otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		return
	}
	otpVal, err := services.EncodeOtpValue(url, user.Code, services.OtpEmail, nil)
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
	return
}
