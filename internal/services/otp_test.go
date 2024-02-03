package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
)

func TestOtpFormatKey(t *testing.T) {
	val, err := services.EncodeOtpValue("url", "code", services.OtpLogin)
	if err != nil {
		t.Error(err)
	}
	if val == "" {
		t.Error("Value is empty")
	}
}
