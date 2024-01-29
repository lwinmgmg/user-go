package controller

import (
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
)

// User Login
func (ctrl *Controller) Login(username, password string, user *models.User) error {
	if err := user.Authenticate(ctrl.RoDb, username, password); err != nil {
		return err
	}
	if user.OtpUrl == "" {
		return nil
	}
	// Otp Authentication is required
	uuid4, err := hashing.NewUuid4Hash256()
	if err != nil {
		return err
	}
	tknExpTime := 5 * time.Minute
	otpVal := services.FormatOtpKey(user.OtpUrl, user.Code, services.OtpLogin)
	if err := ctrl.RedisCtrl.SetKey(string(uuid4), otpVal, tknExpTime); err != nil {
		return err
	}
	// No need to send email for Authenticator User
	if user.IsAuthenticator {
		return nil
	}
	// Need to send email for Non Authenticator User
	partner, err := user.GetPartnerByCode(user.Code, ctrl.RoDb)
	if err != nil {
		return err
	}
	passCode, err := ctrl.Otp.GenerateCode(user.OtpUrl)
	if err != nil {
		return err
	}
	return ctrl.LoginMail.Send(passCode, []string{partner.Email})
}
