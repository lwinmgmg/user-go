package oauth_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestClientScopeTableName(t *testing.T) {
	tbl := oauth.ClientScope{}
	expectedName := models.ComputeTableName("client_scope")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}

func TestCheckClientScope(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		client, _, err := createTestClient(tx)
		if err != nil {
			t.Errorf("Error on creating test client : %v", err)
			return err
		}
		scopes, err := oauth.CheckClientScope(client.ID, tx, "Unknown")
		if err != nil {
			t.Errorf("Error on checking client scope : %v", err)
			return err
		}
		assert.Equal(t, 0, len(scopes), "Expected 0 count in scope : getting %v", len(scopes))
		if err := tx.Model(&oauth.Scope{}).Create(&oauth.Scope{
			Name:  "ExtraUnknown",
			Level: oauth.SL1,
		}).Error; err != nil {
			t.Errorf("Error on creating ExtraUnknown : %v", err)
		}
		scopes, err = oauth.CheckClientScope(client.ID, tx, TestScopeName)
		if err != nil {
			t.Errorf("Error on checking client scope : %v", err)
			return err
		}
		assert.Equal(t, 1, len(scopes), "Expected 1 count in scope : getting %v", len(scopes))
		scopes, err = oauth.CheckClientScope(client.ID, tx, TestScopeName, "ExtraUnknown")
		if !errors.Is(err, oauth.ErrScopeIsNotValid) {
			t.Errorf("Error on checking client scope : %v", err)
		}
		assert.Equal(t, 2, len(scopes), "Expected 2 count in scope : getting %v", len(scopes))
		return errors.New("to_roll_back")
	})
}
