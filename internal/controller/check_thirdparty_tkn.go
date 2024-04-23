package controller

import (
	"encoding/json"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

// Here is the doc string
func (ctrl *Controller) CheckThirdPartyTkn(tkn string) (tknSubject jwtctrl.ThirdPartySubject, err error) {
	if subStr, err := ctrl.RedisCtrl.GetKey(services.FormatThirdpartyTkn(tkn)); err == nil {
		err := json.Unmarshal([]byte(subStr), &tknSubject)
		return tknSubject, err
	}
	_, err = ctrl.JwtCtrl.Validate(tkn, func(c jwt.Claims, t *jwt.Token) (any, error) {
		subStr, err := c.GetSubject()
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal([]byte(subStr), &tknSubject); err != nil {
			return nil, err
		}
		client, err := oauth.GetClientByCId(tknSubject.ClientID, ctrl.RoDb)
		if err != nil {
			return nil, err
		}
		user := models.User{}
		if err := user.GetUserByCode(tknSubject.UserID, ctrl.RoDb); err != nil {
			return nil, err
		}
		ac, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, ctrl.RoDb)
		if err != nil {
			return nil, err
		}
		scopes, err := oauth.GetAcsByUserClientId(ac.ID, ctrl.RoDb)
		if err != nil {
			return nil, err
		}
		expTime, err := c.GetExpirationTime()
		nowTime := time.Now().UTC()
		if err != nil {
			return nil, err
		}
		if nowTime.After(expTime.UTC()) {
			return nil, jwt.ErrTokenExpired
		}
		if err := ctrl.RedisCtrl.SetKey(services.FormatThirdpartyTkn(tkn), subStr, expTime.UTC().Sub(nowTime)); err != nil {
			return nil, err
		}
		for _, scope := range scopes {
			if !slices.Contains(tknSubject.Scopes, scope.Name) {
				return nil, oauth.ErrScopeIsNotValid
			}
		}
		// Get formatted key
		formattedKey := services.FormatJwtKey(client.ClientID, user.Code, client.Secret, ctrl.Setting.JwtService.Key)
		return []byte(formattedKey), nil
	})
	return
}
