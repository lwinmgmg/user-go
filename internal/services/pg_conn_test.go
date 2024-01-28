package services_test

import (
	"testing"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/services"
)

func TestPgConnection(t *testing.T) {
	settings, err := env.LoadSettings()
	if err != nil {
		t.Errorf("Error on getting settings : %v", err)
	}
	conn, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Errorf("Error on psql connection : %v", err)
	}
	if err := conn.Exec("SELECT 1;").Error; err != nil {
		t.Errorf("Error on psql query : %v", err)
	}
}
