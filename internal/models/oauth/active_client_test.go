package oauth_test

import (
	"errors"
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"gorm.io/gorm"
)

func TestActiveClientTableName(t *testing.T) {
	tbl := oauth.ActiveClient{}
	expectedName := models.ComputeTableName("active_client")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}

func TestGetActiveClients(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on getting settings : %v", err)
	}
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		if _, err := oauth.GetActiveClients(1, tx); err != nil {
			t.Errorf("Error on getting acs : %v", err)
		}
		return errors.New("to_roll_back")
	})
}
