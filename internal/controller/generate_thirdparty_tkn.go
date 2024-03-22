package controller

import (
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
)

type ThirdpartyTokenRequest struct {
	Code        string   `form:"code"`
	ClientID    string   `form:"cid"`
	Scopes      []string `form:"scp"`
	RedirectUrl string   `form:"rurl"`
}

func (ctrl *Controller) GenerateThirdPartyToken(userCode string, req ThirdpartyTokenRequest) (loginTkn LoginToken, err error) {
	loginTkn.TokenType = BEARER
	user := models.User{}
	if err = user.GetUserByCode(userCode, ctrl.RoDb); err != nil {
		return
	}
	_, err = oauth.GetClientByCId(req.ClientID, ctrl.RoDb)
	if err != nil {
		return
	}
	return
}
