package services

import (
	"fmt"
	"time"

	otpctrl "github.com/lwinmgmg/user-go/pkg/otp-ctrl"
)

type OtpConfirmType string

const (
	OtpKeyDivider                = "|||"
	OtpLogin      OtpConfirmType = "1"
	OtpEmail      OtpConfirmType = "2"
	OtpPhone      OtpConfirmType = "3"
	OtpAuthr      OtpConfirmType = "4"
	OtpEnable     OtpConfirmType = "5"
)

var (
	OPT_UUID_FORMAT string = fmt.Sprintf("%v%v%v%v%v", "%v", OtpKeyDivider, "%v", OtpKeyDivider, "%v") //otp url, username, type
)

func FormatOtpKey(otpUrl, code string, optType OtpConfirmType) string {
	return fmt.Sprintf(OPT_UUID_FORMAT, otpUrl, code, optType)
}

type OtpService struct {
	otpctrl.OtpCtrl
	Period time.Duration
	Skew   uint
}

func (otpService *OtpService) Validate(passCode, url string) bool {
	return otpService.OtpCtrl.ValidateWithUrl(passCode, url, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC(), otpService.Skew)
}

func (otpService *OtpService) GenerateCode(url string) (string, error) {
	return otpService.OtpCtrl.GenerateCode(url, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC(), otpService.Skew)
}
