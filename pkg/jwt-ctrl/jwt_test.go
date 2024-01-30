package jwtctrl_test

import (
	"testing"
	"time"

	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

func TestJwtCtrl(t *testing.T) {
	jwtC := jwtctrl.NewJwtCtrl("user")
	jwtPassword := "password"
	_, err := jwtC.GenerateCode("lmm", jwtPassword, time.Second)
	if err != nil {
		t.Errorf("Error on generate jwt token : %v", err)
	}
}
