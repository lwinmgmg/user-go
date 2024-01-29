package otpctrl

import (
	"time"

	"github.com/lwinmgmg/user-go/pkg/maths"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

var (
	STANDARD_OPT_DURATION = time.Second * 30
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

func (otpCtrl *OtpCtrl) GenerateCode(url string, duration time.Duration, currentTime time.Time, skews ...uint) (string, error) {
	var skew uint = 0
	key, err := otp.NewKeyFromURL(url)
	if err != nil {
		return "", err
	}
	return totp.GenerateCodeCustom(key.Secret(), currentTime, totp.ValidateOpts{
		Period:    uint(duration.Seconds()),
		Skew:      skew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
}

func (otpCtrl *OtpCtrl) ValidateWithUrl(passcode, url string, duration time.Duration, currentTime time.Time, skews ...uint) bool {
	var skew uint = 0
	if skews != nil {
		skew = maths.Sum(skews)
	}
	key, err := otp.NewKeyFromURL(url)
	if err != nil {
		return false
	}
	res, err := totp.ValidateCustom(
		passcode,
		key.Secret(),
		currentTime,
		totp.ValidateOpts{
			Period:    uint(duration.Seconds()),
			Skew:      skew,
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
