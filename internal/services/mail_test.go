package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/services"
)

func TestNewMailService(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on getting settings : %v", err)
	}
	disableServer := settings.LoginEmailServer
	disableServer.Enable = false
	svr := services.NewMailService(disableServer)
	if err := svr.Send("", []string{""}); err != nil {
		t.Errorf("Getting error on disable email : %v", err)
	}
}
