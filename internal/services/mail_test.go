package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
)

func TestNewMailService(t *testing.T) {
	settings := test.GetTestEnv()
	disableServer := settings.LoginEmailServer
	disableServer.Enable = false
	svr := services.NewMailService(disableServer)
	if err := svr.Send("", []string{""}); err != nil {
		t.Errorf("Getting error on disable email : %v", err)
	}
}
