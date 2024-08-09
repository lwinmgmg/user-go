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

func TestScopeTableName(t *testing.T) {
	tbl := oauth.Scope{}
	expectedName := models.ComputeTableName("scope")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}

func TestGetScopesByClientTID(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(
		func(tx *gorm.DB) error {
			client, _, err := createTestClient(tx)
			if err != nil {
				t.Errorf("Getting error on creating test client : %v", err)
				return err
			}
			scopes, err := oauth.GetScopesByClientTID(client.ID, tx)
			if err != nil {
				t.Errorf("Getting error on get scopes by cid : %v", err)
				return err
			}
			assert.Equalf(t, len(scopes), 1, "Expected length 1 : get %v", len(scopes))
			assert.Equalf(t, scopes[0].Name, "ReadUser", "Scope are not equal : %v", scopes[0].Name)
			return errors.New("to_roll_back")
		})

}
