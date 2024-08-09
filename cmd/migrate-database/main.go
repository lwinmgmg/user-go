package main

import (
	migrateme "github.com/lwinmgmg/user-go/cmd/migrate-database/migrate-me"
	"github.com/lwinmgmg/user-go/env"
)

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	if err := migrateme.MigrateData(settings); err != nil {
		panic(err)
	}
}
