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

func TestActiveClientScopeTableName(t *testing.T) {
	tbl := oauth.ActiveClientScope{}
	expectedName := models.ComputeTableName("active_client_scope")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}

// Test cases for GetAcsByUserId function
func TestGetAcsByUserClientId(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		client, user, err := createTestClient(tx)
		if err != nil {
			t.Errorf("Error on creating test client and user")
			return err
		}
		code := hashing.NewUuid4()
		ac, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, code, tx)
		if err != nil {
			t.Errorf("Error on getting active client : %v", err)
			return err
		}
		acs, err := oauth.GetAcsByUserClientId(ac.ID, tx)
		if len(acs) == 0 && !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("Expected Record Not Found Error : getting %v", err)
		}
		if err := oauth.CreateActiveClientScope(ac.ID, tx, TestScopeName); err != nil {
			t.Errorf("Error on creating Active Client Scope : %v", err)
			return err
		}
		acs, err = oauth.GetAcsByUserClientId(ac.ID, tx)
		if err != nil {
			t.Errorf("Error on getting acs by user client id : %v", err)
			return err
		}
		assert.Equalf(t, 1, len(acs), "Expected scope count : 1, getting : %v", len(acs))
		return errors.New("to_roll_back")
	})
}

// Test case for CreateActiveClientScope
func TestCreateActiveClientScope(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		client, user, err := createTestClient(tx)
		if err != nil {
			t.Errorf("Error on creating test client : %v", err)
			return err
		}
		code := hashing.NewUuid4()
		ac, err := oauth.GetActiveClientCreateIfNotExist(user.ID, client.ID, code, tx)
		if err != nil {
			t.Errorf("Error on creating active client : %v", ac)
			return err
		}
		err = oauth.CreateActiveClientScope(ac.ID, tx, TestScopeName)
		if err != nil {
			return err
		}
		return errors.New("to_roll_back")
	})
}
