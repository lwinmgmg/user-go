package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
)

func TestPgConnection(t *testing.T) {
	settings := test.GetTestEnv()
	conn, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Errorf("Error on psql connection : %v", err)
	}
	if err := conn.Exec("SELECT 1;").Error; err != nil {
		t.Errorf("Error on psql query : %v", err)
	}
}
