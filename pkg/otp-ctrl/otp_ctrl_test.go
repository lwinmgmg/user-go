package otpctrl_test

import (
	"testing"
	"time"

	otpctrl "github.com/lwinmgmg/user-go/pkg/otp-ctrl"
)

func TestOtpCtrl(t *testing.T) {
	ctrl := otpctrl.NewOtpCtrl("user")
	if ctrl.Issuer != "user" {
		t.Errorf("Issuer name miss match : %v", ctrl.Issuer)
	}
	secUrl, err := ctrl.GenerateOtpUrl("abc@gmail.com", otpctrl.STANDARD_OPT_DURATION)
	if err != nil {
		t.Errorf("Error on generate sec otpurl : %v", secUrl)
	}
	code, err := ctrl.GenerateCode(secUrl, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC())
	if err != nil {
		t.Errorf("Error on generating sec code : %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Default Otp code must be 6 digits : %v", code)
	}
	// Unsuccess case
	if ctrl.ValidateWithUrl(code, secUrl, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC().Add(time.Second*61), 1) {
		t.Errorf("Token Should be expired : %v", code)
	}
	// Success case
	newCode, _ := ctrl.GenerateCode(secUrl, otpctrl.STANDARD_OPT_DURATION, time.Now().UTC())
	if !ctrl.ValidateWithUrl(newCode, secUrl, otpctrl.STANDARD_OPT_DURATION, time.Now().Add(time.Second*150).UTC(), 5) {
		t.Errorf("Token Should be valid : %v", newCode)
	}
}
