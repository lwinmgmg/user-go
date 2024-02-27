package controller

import (
	"errors"
	"slices"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
)

var (
	ErrPasswordNotMatch = errors.New("password_not_match")
)

type PasswordChange struct {
	OldPassword string `form:"old_pass"`
	NewPassword string `form:"new_pass"`
}

func (ctrl *Controller) ChangePassword(userCode string, data *PasswordChange) (loginTkn LoginToken, err error) {
	loginTkn.TokenType = OTP_TKN
	loginTkn.SendOtpType = SOtpChPass
	loginTkn.UserCode = userCode
	user := models.User{}
	_, err = user.GetPartnerByCode(userCode, ctrl.RoDb)
	if err != nil {
		return
	}
	var oldPwd []byte
	oldPwd, err = hashing.Hash256(data.OldPassword)
	if err != nil {
		return
	}
	if slices.Compare(oldPwd, user.Password) != 0 {
		err = ErrPasswordNotMatch
		return
	}
	value := map[string]any{
		"password": data.NewPassword,
	}
	if user.IsAuthenticator {
		_, err = GenerateOtp(&loginTkn, user.OtpUrl, user.Code, services.OtpChangePass, ctrl.RedisCtrl, ctrl.Otp,
			time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, value)
		return
	} else {
		var passCode string
		passCode, err = GenerateOtp(&loginTkn, user.OtpUrl, user.Code, services.OtpChangePass, ctrl.RedisCtrl, ctrl.Otp,
			time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, value)
		if err != nil {
			return
		}
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
	}
	err = ErrEmailConfirmRequired
	return
}
