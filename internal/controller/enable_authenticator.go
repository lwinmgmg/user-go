package controller

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image/png"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"github.com/pquerna/otp"
)

var (
	Err2FARequired = errors.New("2fa_required")
)

type Authenticator struct {
	AccessToken string      `json:"access_token"`
	TokenType   TKN_TYPE    `json:"token_type"`
	Image       string      `json:"image"`
	Key         string      `json:"key"`
	UserCode    string      `json:"user_id"`
	SendOtpType SendOtpType `json:"sotp_type"`
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
	uuid4 := hashing.NewUuid4()
	authr.AccessToken = uuid4
	tknExpTime := time.Duration(ctrl.Setting.OtpService.OtpDuration) * time.Second
	otpVal, err := services.EncodeOtpValue(user.OtpUrl, user.Code, services.OtpAuthr, nil)
	if err != nil {
		return
	}
	if err = ctrl.RedisCtrl.SetKey(fmt.Sprintf("otp:%v", uuid4), otpVal, tknExpTime); err != nil {
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
