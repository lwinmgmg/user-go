package models_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
)

func TestUserTableName(t *testing.T) {
	user := models.User{}
	if user.TableName() == "user" {
		t.Errorf("Table prefix is missing for user")
	}
}
