package oauth_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

const (
	TestScopeName = "ReadUser"
)

func TestClientTableName(t *testing.T) {
	client := oauth.Client{}
	if client.TableName() != models.ComputeTableName("client") {
		t.Errorf("Expected %v, getting %v", client.TableName(), models.ComputeTableName("client"))
	}
}

func createTestClient(tx *gorm.DB) (*oauth.Client, *models.User, error) {
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
		Name:        TestScopeName,
		Description: "Read the user info",
		Level:       oauth.SL1,
	}
	if err := tx.Create(&scope).Error; err != nil {
		return nil, nil, err
	}
	if err := tx.Create(&oauth.ClientScope{
		ClientID: client.ID,
		ScopeID:  scope.ID,
	}).Error; err != nil {
		return nil, nil, err
	}
	return &client, user, err
}

func TestGetClientByCid(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Error(err.Error())
	}
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(
		func(tx *gorm.DB) error {
			uuid := hashing.NewUuid4()
			_, err := oauth.GetClientByCId(uuid, tx)
			if err != gorm.ErrRecordNotFound {
				t.Errorf("Expected Record not found error and getting : %v", err)
			}
			client, _, err := createTestClient(tx)
			if err != nil {
				t.Errorf("Error on creating testing client : %v", err)
			}
			client1, err := oauth.GetClientByCId(client.ClientID, tx)
			if err != nil {
				t.Errorf("Error on getting by cid : %v", err)
			}
			if client.ID != client1.ID {
				t.Errorf("Expected : %v, Getting : %v", client.ID, client1.ID)
			}
			return errors.New("to_roll_back")
		},
	)
}
