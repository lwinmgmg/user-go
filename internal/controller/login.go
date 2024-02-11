package controller

import (
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
)

type TKN_TYPE string

var (
	BEARER  TKN_TYPE = "Bearer"
	OTP_TKN TKN_TYPE = "Otp"
)

type LoginToken struct {
	TokenType   TKN_TYPE `json:"token_type"`
	AccessToken string   `json:"access_token"`
	UserCode    string   `json:"user_id"`
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
	uuid4 := hashing.NewUuid4() + hashing.NewUuid4() + username
	loginTkn.TokenType = OTP_TKN
	loginTkn.AccessToken = string(uuid4)
	tknExpTime := 5 * time.Minute
	otpVal, err := services.EncodeOtpValue(user.OtpUrl, user.Code, services.OtpLogin)
	if err != nil {
		return loginTkn, err
	}
	if err := ctrl.RedisCtrl.SetKey(fmt.Sprintf("otp:%v", uuid4), otpVal, tknExpTime); err != nil {
		return loginTkn, err
	}
	// No need to send email for Authenticator User
	if user.IsAuthenticator {
		return loginTkn, nil
	}
	// Need to send email for Non Authenticator User
	partner, err := user.GetPartnerByCode(user.Code, ctrl.RoDb)
	if err != nil {
		return loginTkn, err
	}
	passCode, err := ctrl.Otp.GenerateCode(user.OtpUrl)
	if err != nil {
		return loginTkn, err
	}
	if user.Partner.IsEmailConfirmed {
		go ctrl.LoginMail.Send(passCode, []string{partner.Email})
	} else if user.Partner.IsPhoneConfirmed {
		go ctrl.PhoneService.Send(passCode, user.Partner.Phone)
	}
	return loginTkn, nil
}
