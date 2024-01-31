package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

var (
	ErrEmptyAuth   = errors.New("empty_authorization")
	ErrInvalidAuth = errors.New("invalid_authorization")
)

func parseToken(str string) (string, error) {
	r, err := regexp.Compile("^Bearer (.+)")
	if err != nil {
		return "", err
	}
	strs := r.FindStringSubmatch(str)
	if len(strs) != 2 {
		return "", ErrInvalidAuth
	}
	return strs[1], nil
}

func (apiCtrl *ApiCtrl) AuthMiddleware(ctx *gin.Context) {
	value := ctx.Request.Header.Get("Authorization")
	if value == "" {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 1, "Authorization Required!", ErrEmptyAuth))
	}
	token, err := parseToken(value)
	if err != nil {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 2, "Authorization Required!", err))
	}
	sub := &jwtctrl.Subject{}
	if _, err := apiCtrl.JwtCtrl.Validate(token, func(c jwt.Claims, t *jwt.Token) (any, error) {
		subStr, err := c.GetSubject()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(subStr), sub); err != nil {
			return nil, err
		}
		return []byte(apiCtrl.Settings.JwtService.Key), nil
	}); err != nil {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 3, "Authorization Required!", err))
	}
	ctx.Set("subject", sub)
}
