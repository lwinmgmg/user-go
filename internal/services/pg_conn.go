package services

import (
	"fmt"
	"net/url"
	"os"

	"github.com/lwinmgmg/user-go/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPsql(dbConf *env.DbServer) (*gorm.DB, error) {
	var dsn string = fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable TimeZone=UTC",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		url.QueryEscape(dbConf.Password),
		dbConf.DbName,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	val, _ := os.LookupEnv("GIN_MODE")
	if val != "release" {
		return db.Debug(), err
	}
	return db, err
}
