package main

import (
	"fmt"

	migrateme "github.com/lwinmgmg/user-go/cmd/migrate-database/migrate-me"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"gorm.io/gorm"
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
	ac := &oauth.ActiveClient{}
	acs := &oauth.ActiveClientScope{}
	if err := db.Transaction(
		func(tx *gorm.DB) error {
			if err := db.Migrator().AutoMigrate(
				user,
				partner,
				client,
				scope,
				cs,
				ac,
				acs,
			); err != nil {
				return err
			}
			if err := db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v;", partner.GetSequence())).Error; err != nil {
				return err
			}
			if err := db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v START WITH 100000;", user.GetSequence())).Error; err != nil {
				return err
			}
			return migrateme.MigrateDefaultData(&settings, tx)
		},
	); err != nil {
		panic(err)
	}
}
