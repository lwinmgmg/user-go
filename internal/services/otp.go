package services

import (
	"encoding/json"
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
	OtpChangePass OtpConfirmType = "6"
)

var (
	OPT_UUID_FORMAT string = fmt.Sprintf("%v%v%v%v%v", "%v", OtpKeyDivider, "%v", OtpKeyDivider, "%v") //otp url, username, type
)

type OtpValue struct {
	Url   string         `json:"url"`
	Code  string         `json:"code"`
	Type  OtpConfirmType `json:"type"`
	Value map[string]any `json:"value"`
}

func EncodeOtpValue(otpUrl, code string, otpType OtpConfirmType, value map[string]any) (string, error) {
	otpValue := OtpValue{
		Url:  otpUrl,
		Code: code,
		Type: otpType,
	}
	val, err := json.Marshal(otpValue)
	return string(val), err
}

func ParseOtpValue(otpStr string) (OtpValue, error) {
	otpValue := OtpValue{}
	if err := json.Unmarshal([]byte(otpStr), &otpValue); err != nil {
		return otpValue, err
	}
	return otpValue, nil
}

type OtpService struct {
	*otpctrl.OtpCtrl
	Period time.Duration
	Skew   uint
}

func (otpService *OtpService) Validate(passCode, url string) bool {
	return otpService.OtpCtrl.ValidateWithUrl(passCode, url, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC(), otpService.Skew)
}

func (otpService *OtpService) GenerateCode(url string) (string, error) {
	return otpService.OtpCtrl.GenerateCode(url, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC(), otpService.Skew)
}

func NewOtpService(otpCtrl *otpctrl.OtpCtrl, period time.Duration, skew uint) *OtpService {
	return &OtpService{
		OtpCtrl: otpCtrl,
		Period:  period,
		Skew:    skew,
	}
}
