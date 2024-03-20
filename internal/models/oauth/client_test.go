package oauth_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
)

func TestClientTableName(t *testing.T) {
	client := oauth.Client{}
	if client.TableName() != models.ComputeTableName("client") {
		t.Errorf("Expected %v, getting %v", client.TableName(), models.ComputeTableName("client"))
	}
}
