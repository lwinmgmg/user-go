package test

import (
	"fmt"

	migrateme "github.com/lwinmgmg/user-go/cmd/migrate-database/migrate-me"
	"github.com/lwinmgmg/user-go/env"
)

var TestUserDBName string = "user_go_test"

func init() {
	settings := GetTestEnv()
	migrateme.MigrateData(settings)
}

func GetTestEnv() *env.Settings {
	settings, err := env.LoadSettings()
	if err != nil {
		fmt.Println("Error on loading test env")
		panic(err)
	}
	settings.Db.DbName = TestUserDBName
	settings.RoDb.DbName = TestUserDBName
	return settings
}
