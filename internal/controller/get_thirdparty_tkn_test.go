package controller_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestGetThirdPartyToken(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	ctrl := controller.NewContoller(settings)
	scopes := []string{"ReadUser"}
	db.Transaction(func(tx *gorm.DB) error {
		client, user, err := createTestClientUser(tx)
		if err != nil {
			return err
		}
		ctrl.RoDb = tx
		ctrl.Db = tx
		oldTkn, code, err := ctrl.GenerateThirdPartyToken(user.Code, client.ClientID, client.RedirectUrl, scopes...)
		if err != nil {
			t.Error("Error on generating thirdparty tkn", err)
			return err
		}
		tkn, err := ctrl.GetThirdPartyToken(code, client.ClientID, client.Secret)
		if err != nil {
			t.Error("Error getting third party token", err)
			return err
		}
		assert.Equal(t, oldTkn.AccessToken, tkn.AccessToken)
		assert.Equal(t, oldTkn.SendOtpType, tkn.SendOtpType)
		assert.Equal(t, oldTkn.TokenType, tkn.TokenType)
		assert.Equal(t, oldTkn.UserCode, tkn.UserCode)
		return errors.New("to_roll_back")
	})
}
