package models

import (
	"errors"
	"time"

	"github.com/lwinmgmg/user-go/env"
	uuid_code "github.com/lwinmgmg/uuid_code/v1"
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
