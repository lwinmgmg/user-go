package controller_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createTestClientUser(tx *gorm.DB) (*oauth.Client, *models.User, error) {
	uuid := hashing.NewUuid4()
	secret := hashing.NewUuid4()
	user, err := models.CreateTestUser("testing", "testing", tx)
	if err != nil {
		return nil, nil, err
	}
	client := oauth.Client{
		Name:          "testing",
		ClientID:      uuid,
		Secret:        secret,
		UserID:        user.ID,
		RedirectUrl:   "http://localhost",
		VerifiedLevel: oauth.SL1,
	}
	if err := tx.Create(&client).Error; err != nil {
		return nil, nil, err
	}
	scope := oauth.Scope{
		Name:        "ReadUser",
		Description: "Read the user info",
		Level:       oauth.SL1,
	}
	if err := tx.Create(&scope).Error; err != nil {
		return nil, nil, err
	}
	return &client, user, err
}
func TestGenerateThirdPartyTkn(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on getting settings : %v", err)
	}
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	ctrl := controller.NewContoller(settings)
	db.Transaction(func(tx *gorm.DB) error {
		ctrl.RoDb = tx
		client, user, err := createTestClientUser(tx)
		if err != nil {
			t.Errorf("Error on creating client,user : %v", err)
		}
		tptyReq := controller.ThirdpartyTokenRequest{
			Code:        "12345",
			ClientID:    client.ClientID,
			Scopes:      []string{"ReadUser"},
			RedirectUrl: "http://localhost",
		}
		loginTkn, err := ctrl.GenerateThirdPartyToken("no_user", tptyReq)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("Expected %v, getting %v", gorm.ErrRecordNotFound, err)
		}
		loginTkn, err = ctrl.GenerateThirdPartyToken(user.Code, tptyReq)
		if err != nil {
			t.Errorf("Error on generating third party token, %v", err)
			return err
		}
		if loginTkn.TokenType != controller.BEARER {
			assert.Equal(t, loginTkn.TokenType, controller.BEARER, "Expected Bearer Token")
		}
		return errors.New("to_roll_back")
	})
}
