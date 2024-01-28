package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
)

func TestOtpFormatKey(t *testing.T) {
	val := services.FormatOtpKey("url", "code", services.OtpLogin)
	if val == "" {
		t.Error("Value is empty")
	}
}
