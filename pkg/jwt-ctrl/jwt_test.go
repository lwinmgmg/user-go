package jwtctrl_test

import (
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

func TestJwtCtrl(t *testing.T) {
	jwtC := jwtctrl.NewJwtCtrl("user")
	jwtPassword := "password"
	tkn, err := jwtC.GenerateCode("lmm", jwtPassword, time.Second)
	if err != nil {
		t.Errorf("Error on generate jwt token : %v", err)
	}
	if _, err := jwtC.Validate(tkn, func(c jwt.Claims, t *jwt.Token) (any, error) {
		return []byte(jwtPassword), err
	}); err != nil {
		t.Errorf("Error on validating Token : %v", err)
	}
	tkn1, err := jwtC.GenerateCode("lmm", jwtPassword, time.Nanosecond)
	if err != nil {
		t.Errorf("error on generate jwt token with nano : %v", err)
	}
	if _, err := jwtC.Validate(tkn1, func(c jwt.Claims, t *jwt.Token) (any, error) {
		return []byte(jwtPassword), nil
	}); !errors.Is(err, jwt.ErrTokenExpired) {
		t.Errorf("Expected token expired, getting : %v", err)
	}
}
