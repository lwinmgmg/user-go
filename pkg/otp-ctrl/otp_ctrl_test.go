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
	secUrl, err := ctrl.GenerateOtpUrl("abc@gmail.com", time.Second*30)
	if err != nil {
		t.Errorf("Error on generate sec otpurl : %v", secUrl)
	}
	code, err := ctrl.GenerateCode(secUrl)
	if err != nil {
		t.Errorf("Error on generating sec code : %v", err)
	}
	if len(code) != 6 {
		t.Errorf("Default Otp code must be 6 digits : %v", code)
	}
	// Unsuccess case
	if ctrl.ValidateWithUrl(code, secUrl, time.Second) {
		t.Errorf("Token Should be expired : %v", code)
	}
	// Success case
	newCode, _ := ctrl.GenerateCode(secUrl)
	if !ctrl.ValidateWithUrl(newCode, secUrl, time.Second) {
		t.Errorf("Token Should be valid : %v", newCode)
	}
}
