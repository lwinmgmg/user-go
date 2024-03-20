package oauth_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
)

func TestClientScopeTableName(t *testing.T) {
	tbl := oauth.ClientScope{}
	expectedName := models.ComputeTableName("client_scope")
	if tbl.TableName() != expectedName {
		t.Errorf("Expected %v, getting %v", expectedName, tbl.TableName())
	}
}
