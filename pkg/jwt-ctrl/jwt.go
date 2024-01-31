package jwtctrl

import (
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Subject struct {
	ClientID  string `json:"cid"`
	UserID    string `json:"uid"`
	Reference string `json:"ref"`
}

type JwtCtrl struct {
	Issuer string
}

func (jwtCtrl *JwtCtrl) GenerateCode(subject Subject, key string, duration time.Duration, audiences ...string) (string, error) {
	nowTime := time.Now().UTC()
	subStr, err := json.Marshal(subject)
	if err != nil {
		return "", err
	}
	claim := jwt.RegisteredClaims{
		Issuer:    jwtCtrl.Issuer,
		IssuedAt:  jwt.NewNumericDate(nowTime),
		ExpiresAt: jwt.NewNumericDate(nowTime.Add(duration)),
		Subject:   string(subStr),
		Audience:  audiences,
	}
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return tkn.SignedString([]byte(key))
}

func (JwtCtrl *JwtCtrl) Validate(tkn string, getKey func(jwt.Claims, *jwt.Token) (any, error)) (jwt.Claims, error) {
	claim := &jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tkn, claim, func(t *jwt.Token) (interface{}, error) {
		return getKey(claim, t)
	})
	return claim, err
}

func NewJwtCtrl(issuer string) *JwtCtrl {
	return &JwtCtrl{
		Issuer: issuer,
	}
}
