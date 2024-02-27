package controller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image/png"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/pquerna/otp"
)

var (
	Err2FARequired = errors.New("2fa_required")
)

type Authenticator struct {
	LoginToken
	Image string `json:"image"`
	Key   string `json:"key"`
}

func (ctrl *Controller) EnableAuthenticator(userCode string) (authr Authenticator, err error) {
	authr.TokenType = OTP_TKN
	authr.UserCode = userCode
	authr.SendOtpType = SOtpAuth
	user := models.User{}
	if _, err = user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	if user.OtpUrl == "" {
		err = Err2FARequired
		return
	}
	_, err = GenerateOtp(&authr.LoginToken, user.OtpUrl, user.Code, services.OtpAuthr, ctrl.RedisCtrl, ctrl.Otp,
		time.Duration(ctrl.Setting.OtpService.OtpDuration)*time.Second, nil)
	if err != nil {
		return
	}
	key, err := otp.NewKeyFromURL(user.OtpUrl)
	if err != nil {
		return
	}
	authr.Key = key.Secret()
	img, err := key.Image(100, 100)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	if err = png.Encode(&buf, img); err != nil {
		return
	}
	authr.Image = base64.StdEncoding.EncodeToString(buf.Bytes())
	return
}
