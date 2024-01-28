package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/services"
)

func TestRedisConn(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on getting env : %v", err)
	}
	_, err = services.GetRedisClient(settings.Redis)
	if err != nil {
		t.Errorf("Error on getting redis : %v", err)
	}
}
