package controller

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

var (
	ErrInvalidThirdParty error = errors.New("ErrInvalidThirdParty")
)

func (ctrl *Controller) GetThirdPartyToken(code, userCode, clientId, clientSecret string) (loginTkn LoginToken, err error) {
	key := services.FormatThirdpartyCode(code, clientId)
	tknStr, err := ctrl.RedisCtrl.GetKey(key, time.Second*5)
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(tknStr), &loginTkn); err != nil {
		return
	}
	thirdPartySub := &jwtctrl.ThirdPartySubject{}
	if err = services.ParseJwtToken(thirdPartySub, loginTkn.AccessToken, services.FormatJwtKey(clientId, userCode, clientSecret, ctrl.Setting.JwtService.Key), ctrl.JwtCtrl); err != nil {
		return
	}
	client, err := oauth.GetClientByCId(thirdPartySub.ClientID, ctrl.RoDb)
	if err != nil {
		return
	}
	user := models.User{}
	if err = user.GetUserByCode(thirdPartySub.UserID, ctrl.RoDb); err != nil {
		return
	}
	activeClient, err := oauth.GetActiveClient(user.ID, client.ID, ctrl.RoDb)
	if err != nil {
		return
	}
	if client.ClientID != clientId || client.Secret != clientSecret || activeClient.RefreshToken != code {
		err = ErrInvalidThirdParty
		return
	}
	return
}
