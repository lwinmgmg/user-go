package services

import (
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/env"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

func FormatJwtKey(username, userCode, password, key string) string {
	return fmt.Sprintf("%v:%v:%v:%v", username, userCode, password, key)
}

func GenerateUserLoginJwt(userCode, formattedKey string, settings *env.Settings, jwtCtrl *jwtctrl.JwtCtrl) (string, error) {
	return jwtCtrl.GenerateCode(jwtctrl.UserSubject{
		UserID: userCode,
	}, formattedKey, time.Second*time.Duration(settings.JwtService.LoginDuration), settings.Service)
}

func GenerateThirdpartyJwt(userCode, clientId, formattedKey string, settings *env.Settings, jwtCtrl *jwtctrl.JwtCtrl, scopes ...string) (string, error) {
	return jwtCtrl.GenerateCode(
		jwtctrl.ThirdPartySubject{
			ClientID: clientId,
			UserID:   userCode,
			Scopes:   scopes,
		},
		formattedKey, time.Second*time.Duration(settings.JwtService.LoginDuration), settings.Service,
	)
}

func FormatThirdpartyTkn(tkn string) string {
	return fmt.Sprintf("thirdparty:%v", tkn)
}
