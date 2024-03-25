package controller

import (
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
)

func (ctrl *Controller) CheckThirdPartyTkn(tkn string) (err error) {
	_, err = ctrl.JwtCtrl.Validate(tkn, func(c jwt.Claims, t *jwt.Token) (any, error) {
		subStr, err := c.GetSubject()
		if err != nil {
			return nil, err
		}
		sub := &jwtctrl.ThirdPartySubject{}
		if err := json.Unmarshal([]byte(subStr), sub); err != nil {
			return nil, err
		}
		return []byte(""), nil
	})
	return
}
