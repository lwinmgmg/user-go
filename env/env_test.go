package env_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/env"
)

func TestLoadSettings(t *testing.T) {
	_, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on loading setting : %v", err)
	}
}
