package env_test

import (
	"errors"
	"os"
	"testing"

	"github.com/lwinmgmg/user-go/env"
)

func TestLoadSettings(t *testing.T) {
	_, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on loading setting : %v", err)
	}
	if os.Setenv("USER_SETTING_PATH", "env.go") != nil {
		t.Errorf("Error on loading setting : %v", err)
	}
	_, err = env.LoadSettings()
	if err == nil {
		t.Errorf("Expected Error on yaml unmarshal")
	}
	if os.Unsetenv("USER_SETTING_PATH") != nil {
		t.Errorf("Error on loading setting : %v", err)
	}
	_, err = env.LoadSettings()
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("Expected no such file")
	}
}
