package services

import (
	"fmt"
	"net/url"

	"github.com/lwinmgmg/user-go/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPsql(dbConf env.DbServer) (*gorm.DB, error) {
	var dsn string = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=UTC",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		url.QueryEscape(dbConf.Password),
		dbConf.DbName,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
}
