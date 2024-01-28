package controller

import (
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
)

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
	if user.IsAuthenticator {
		return nil
	}
	partner, err := user.GetPartnerByCode(user.Code, ctrl.RoDb)
	if err != nil {
		return err
	}
	ctrl.LoginMail.Send("", []string{partner.Email})
	return nil
}
