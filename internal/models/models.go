package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/lwinmgmg/user-go/env"
	uuid_code "github.com/lwinmgmg/uuid_code/v1"
	"gorm.io/gorm"
)

var (
	uuidCode              *uuid_code.UuidCode = uuid_code.NewDefaultUuidCode()
	Env                   env.Settings
	ErrRecordAlreadyExist = errors.New("exist")
	ErrInvalid            = errors.New("invalid")
)

func init() {
	var err error
	Env, err = env.LoadSettings()
	if err != nil {
		panic(err)
	}
}

func computeTableName(input string) string {
	if Env.Db.TablePrefix == "" {
		return input
	}
	return Env.Db.TablePrefix + "_" + input
}

func ComputeTableName(input string) string {
	return computeTableName(input)
}

type DefaultModel struct {
	ID         uint      `gorm:"primaryKey"`
	CreateDate time.Time `gorm:"autoCreateTime:nano"`
	WriteDate  time.Time `gorm:"autoUpdateTime:nano"`
}

func InitDb(db *gorm.DB) error {
	user := &User{}
	partner := &Partner{}
	if err := db.Migrator().AutoMigrate(
		user,
		partner,
	); err != nil {
		return err
	}
	db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v;", partner.GetSequence()))
	db.Exec(fmt.Sprintf("CREATE SEQUENCE IF NOT EXISTS %v START WITH 100000;", user.GetSequence()))
	return nil
}
