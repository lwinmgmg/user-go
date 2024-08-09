package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
)

func TestRedisConn(t *testing.T) {
	settings := test.GetTestEnv()
	_, err := services.GetRedisClient(settings.Redis)
	if err != nil {
		t.Errorf("Error on getting redis : %v", err)
	}
}
