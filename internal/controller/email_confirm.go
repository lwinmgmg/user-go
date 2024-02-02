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
	loginTkn.TokenType = OTP_TKN
	user := models.User{}
	if _, err = user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	if user.Partner.IsEmailConfirmed {
		err = ErrEmailAldyConfirmed
		return
	}
	if err = user.Partner.CheckEmail(ctrl.RoDb); err != nil {
		return
	}
	uuid4 := hashing.NewUuid4()
	loginTkn.TokenType = OTP_TKN
	loginTkn.Value = string(uuid4)
	tknExpTime := 5 * time.Minute
	otpVal := services.FormatOtpKey(user.OtpUrl, user.Code, services.OtpEmail)
	if err = ctrl.RedisCtrl.SetKey(fmt.Sprintf("otp:%v", string(uuid4)), otpVal, tknExpTime); err != nil {
		return
	}
	url, err := ctrl.Otp.OtpCtrl.GenerateOtpUrl(user.Partner.Email, otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		return
	}
	passCode, err := ctrl.Otp.GenerateCode(url)
	if err != nil {
		return
	}
	err = ctrl.LoginMail.Send(passCode, []string{user.Partner.Email})
	return
}
