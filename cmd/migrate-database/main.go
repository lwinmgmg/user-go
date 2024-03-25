package main

import (
	"fmt"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
)

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		panic(err)
	}
	user := &models.User{}
	partner := &models.Partner{}
	client := &oauth.Client{}
	scope := &oauth.Scope{}
	cs := &oauth.ClientScope{}
	acs := &oauth.ActiveClient{}
	if err := db.Migrator().AutoMigrate(
		user,
		partner,
		client,
		scope,
		cs,
		acs,
	); err != nil {
		panic(err)
	}
	// Creating sequence
	db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v;", partner.GetSequence()))
	db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v START WITH 100000;", user.GetSequence()))
}
