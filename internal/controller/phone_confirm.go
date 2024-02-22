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
	ErrPhoneAldyConfirmed = errors.New("phone_already_confirmed")
)

func (ctrl *Controller) PhoneConfirm(userCode string) (loginTkn LoginToken, err error) {
	loginTkn.SendOtpType = SOtpPhone
	loginTkn.TokenType = OTP_TKN
	user := models.User{}
	if _, err = user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	loginTkn.UserCode = user.Code
	if user.Partner.IsPhoneConfirmed {
		err = ErrPhoneAldyConfirmed
		return
	}
	if err = user.Partner.CheckPhone(ctrl.RoDb); err != nil {
		return
	}
	uuid4 := hashing.NewUuid4()
	loginTkn.AccessToken = string(uuid4)
	tknExpTime := time.Duration(ctrl.Setting.OtpService.OtpDuration) * time.Second
	url, err := ctrl.Otp.OtpCtrl.GenerateOtpUrl(user.Partner.Email, otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		return
	}
	otpVal, err := services.EncodeOtpValue(url, user.Code, services.OtpPhone, nil)
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
	err = ctrl.PhoneService.Send(passCode, user.Partner.Phone)
	return
}
