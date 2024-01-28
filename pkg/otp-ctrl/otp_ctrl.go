package otpctrl

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OtpCtrl struct {
	Issuer string
}

func (otpCtrl *OtpCtrl) GenerateOtpUrl(code string, duration time.Duration) (string, error) {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      otpCtrl.Issuer,
		AccountName: code,
		Period:      uint(duration.Seconds()),
	})
	if err != nil {
		return "", err
	}
	return key.URL(), err
}

func (otpCtrl *OtpCtrl) GenerateCode(url string) (string, error) {
	key, err := otp.NewKeyFromURL(url)
	if err != nil {
		return "", err
	}
	return totp.GenerateCode(key.Secret(), time.Now().UTC())
}

func (otpCtrl *OtpCtrl) ValidateWithUrl(passcode, url string, duration time.Duration) bool {
	key, err := otp.NewKeyFromURL(url)
	if err != nil {
		return false
	}
	res, err := totp.ValidateCustom(
		passcode,
		key.Secret(),
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    uint(duration.Seconds()),
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		})
	if err != nil {
		return false
	}
	return res
}

func NewOtpCtrl(issurer string) *OtpCtrl {
	return &OtpCtrl{
		Issuer: issurer,
	}
}
