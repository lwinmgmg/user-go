package oauth_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
)

func TestScopeTableName(t *testing.T) {
	tbl := oauth.Scope{}
	expectedName := models.ComputeTableName("scope")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}
