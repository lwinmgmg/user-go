package main

import (
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
	if err := models.InitDb(db); err != nil {
		panic(err)
	}
	if err := oauth.InitDb(db); err != nil {
		panic(err)
	}
}
