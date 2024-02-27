package controller

import (
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
)

type TKN_TYPE string

var (
	BEARER  TKN_TYPE = "Bearer"
	OTP_TKN TKN_TYPE = "Otp"
)

type LoginToken struct {
	TokenType   TKN_TYPE    `json:"token_type"`
	AccessToken string      `json:"access_token"`
	UserCode    string      `json:"user_id"`
	SendOtpType SendOtpType `json:"sotp_type"`
}

// User Login
func (ctrl *Controller) Login(username, password string, user *models.User) (*LoginToken, error) {
	loginTkn := &LoginToken{
		TokenType: BEARER,
	}
	if err := user.Authenticate(ctrl.RoDb, username, password); err != nil {
		return loginTkn, err
	}
	loginTkn.UserCode = user.Code
	if user.OtpUrl == "" {
		formattedKey := services.FormatJwtKey(user.Username, user.Code, string(user.Password), ctrl.Setting.JwtService.Key)
		jwtToken, err := services.GenerateUserLoginJwt(user.Code, formattedKey, ctrl.Setting, ctrl.JwtCtrl)
		if err != nil {
			return loginTkn, err
		}
		loginTkn.AccessToken = jwtToken
		return loginTkn, nil
	}
	// Otp Authentication is required
	passCode, err := GenerateOtp(loginTkn, user.OtpUrl, user.Code, services.OtpLogin, ctrl.RedisCtrl, ctrl.Otp,
		time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, nil)
	if err != nil {
		return loginTkn, err
	}
	// No need to send email for Authenticator User
	if user.IsAuthenticator {
		loginTkn.SendOtpType = SOtpAuth
		return loginTkn, nil
	}
	partner, err := user.GetPartnerByCode(user.Code, ctrl.RoDb)
	if err != nil {
		return loginTkn, err
	}
	if user.Partner.IsEmailConfirmed {
		loginTkn.SendOtpType = SOtpEmail
		go ctrl.LoginMail.Send(passCode, []string{partner.Email})
	} else if user.Partner.IsPhoneConfirmed {
		loginTkn.SendOtpType = SOtpPhone
		go ctrl.PhoneService.Send(passCode, user.Partner.Phone)
	}
	return loginTkn, nil
}
