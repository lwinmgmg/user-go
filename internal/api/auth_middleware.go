package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

var (
	ErrEmptyAuth       = errors.New("empty_authorization")
	ErrInvalidAuth     = errors.New("invalid_authorization")
	ErrSubjectNotFound = errors.New("subject_not_found")
)

func DelFormatReditToken(userCode string) string {
	return fmt.Sprintf("del:%v", userCode)
}

func formatRedisToken(token string) string {
	return fmt.Sprintf("token:%v", token)
}

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

func getUserSubject[T jwtctrl.Subject](ctx *gin.Context) *T {
	sub, ok := ctx.Get("subject")
	if !ok {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 6, "Authorization Required!", ErrSubjectNotFound))
	}
	jwtSub, ok := sub.(*T)
	if !ok {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 7, "Authorization Required!", ErrSubjectNotFound))
	}
	return jwtSub
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
	sub := &jwtctrl.UserSubject{}
	if val, err := apiCtrl.RedisCtrl.GetKey(formatRedisToken(token)); err == nil {
		if err := json.Unmarshal([]byte(val), sub); err == nil {
			ctx.Set("subject", sub)
			return
		}
	}
	if _, err := apiCtrl.JwtCtrl.Validate(token, func(c jwt.Claims, t *jwt.Token) (any, error) {
		subStr, err := c.GetSubject()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(subStr), sub); err != nil {
			return nil, err
		}
		user := &models.User{}
		if err := user.GetUserByCode(sub.UserID, apiCtrl.RoDb); err != nil {
			panic(middlewares.NewPanic(http.StatusUnauthorized, 4, "Authorization Required!", err))
		}
		fTkn := formatRedisToken(token)
		if err := apiCtrl.RedisCtrl.SetKey(fTkn, subStr, time.Second*time.Duration(apiCtrl.Settings.JwtService.CacheDuration)); err != nil {
			slog.Error(fmt.Sprintf("Can't set jwt cache in redis %v", err))
		}
		if err := apiCtrl.RedisCtrl.SetKey(DelFormatReditToken(user.Code), fTkn, time.Second*time.Duration(apiCtrl.Settings.JwtService.CacheDuration)); err != nil {
			slog.Error(fmt.Sprintf("Can't set jwt cache in redis %v", err))
		}
		formattedKey := services.FormatJwtKey(user.Username, user.Code, string(user.Password), apiCtrl.Settings.JwtService.Key)
		return []byte(formattedKey), nil
	}); err != nil {
		panic(middlewares.NewPanic(http.StatusUnauthorized, 5, "Authorization Required!", err))
	}
	ctx.Set("subject", sub)
}
