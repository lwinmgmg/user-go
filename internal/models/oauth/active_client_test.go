package oauth_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"github.com/lwinmgmg/user-go/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestActiveClientTableName(t *testing.T) {
	tbl := oauth.ActiveClient{}
	expectedName := models.ComputeTableName("active_client")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}

func TestGetActiveClientCreateIfNotExist(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		client, user, err := createTestClient(tx)
		if err != nil {
			t.Errorf("Getting error on creating test client user : %v", err)
			return err
		}
		code := hashing.NewUuid4()
		ac, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, code, tx)
		if err != nil {
			t.Errorf("Error on creating acs : %v", err)
		}
		if _, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, code, tx); err != nil {
			t.Errorf("Error on getting existing acs : %v", err)
		}
		var count int64 = 0
		if err := tx.Model(&oauth.ActiveClient{}).Where("client_id=? AND user_id=?", client.ID, user.ID).Count(&count).Error; err != nil {
			t.Errorf("Error on getting active client count : %v", err)
		}
		assert.Equalf(t, int64(1), count, "Expected count = 1, getting %v", count)
		assert.Equal(t, ac.RefreshToken, code, "Expected code are equal")
		return errors.New("to_roll_back")
	})
}
