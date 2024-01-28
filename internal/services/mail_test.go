package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
)

func TestNewMailService(t *testing.T) {
	svr := services.NewMailService("", "", "", 443, false)
	if err := svr.Send("", []string{""}); err != nil {
		t.Errorf("Getting error on disable email : %v", err)
	}
}
