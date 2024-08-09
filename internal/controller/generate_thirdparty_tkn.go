package controller

import (
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"gorm.io/gorm"
)

var (
	ErrUnauthorizedScope = errors.New("unauthorized_scope")
	ErrNoScope           = errors.New("no_scope")
	ErrMissMatchRUrl     = errors.New("missmatch_rurl")
)

func (ctrl *Controller) GenerateThirdPartyToken(userCode, clientId, redirectUrl string, inputScopes ...string) (loginTkn LoginToken, code string, err error) {
	loginTkn.TokenType = BEARER
	code = hashing.NewUuid4()
	user := models.User{}
	if err = user.GetUserByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	client, err := oauth.GetClientByCId(clientId, ctrl.RoDb)
	if err != nil {
		return
	}
	defer func() {
		if err == nil {
			err = ctrl.Db.Transaction(
				func(tx *gorm.DB) error {
					ac, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, code, tx)
					if err != nil {
						return err
					}
					subStr, err := json.Marshal(jwtctrl.ThirdPartySubject{
						ClientID: clientId,
						UserID:   userCode,
						Scopes:   inputScopes,
					})
					if err != nil {
						return err
					}
					if err := ctrl.RedisCtrl.SetKey(services.FormatThirdpartyTkn(loginTkn.AccessToken), string(subStr), time.Second*time.Duration(ctrl.Setting.JwtService.LoginDuration)); err != nil {
						return err
					}
					return oauth.CreateActiveClientScope(ac.ID, tx, inputScopes...)
				},
			)
		}
	}()
	if redirectUrl != client.RedirectUrl {
		err = ErrMissMatchRUrl
		return
	}

	scopes, err := oauth.GetScopesByClientTID(client.ID, ctrl.RoDb)
	if err != nil {
		return
	}
	scopeNames := make([]string, len(scopes))
	for i := 0; i < len(scopes); i++ {
		if client.VerifiedLevel >= scopes[i].Level {
			scopeNames = append(scopeNames, scopes[i].Name)
		}
	}
	if len(inputScopes) == 0 {
		err = ErrNoScope
		return
	}
	for i := 0; i < len(inputScopes); i++ {
		if !slices.Contains(scopeNames, inputScopes[i]) {
			err = ErrUnauthorizedScope
			return
		}
	}
	// Get formatted key
	formattedKey := services.FormatJwtKey(client.ClientID, user.Code, client.Secret, ctrl.Setting.JwtService.Key)
	// Generate token
	if loginTkn.AccessToken, err = services.GenerateThirdpartyJwt(user.Code, client.ClientID, formattedKey, ctrl.Setting, ctrl.JwtCtrl, inputScopes...); err != nil {
		return
	}
	loginTkn.UserCode = userCode
	value, err := json.Marshal(loginTkn)
	if err != nil {
		return
	}
	if err = ctrl.RedisCtrl.SetKey(services.FormatThirdpartyCode(code, clientId), string(value), 5*time.Minute); err != nil {
		return
	}
	return
}
