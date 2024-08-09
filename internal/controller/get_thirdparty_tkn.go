package controller

import (
	"encoding/json"
	"time"

	"github.com/lwinmgmg/user-go/internal/services"
)

func (ctrl *Controller) GetThirdPartyToken(code, clientId string, clientSecret string) (loginTkn LoginToken, err error) {
	key := services.FormatThirdpartyCode(code, clientId)
	tknStr, err := ctrl.RedisCtrl.GetKey(key, time.Second*5)
	if err != nil {
		return
	}
	if err = json.Unmarshal([]byte(tknStr), &loginTkn); err != nil {
		return
	}
	return
}
