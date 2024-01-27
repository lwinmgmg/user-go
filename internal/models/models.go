package models

import (
	"errors"
	"time"

	uuid_code "github.com/lwinmgmg/uuid_code/v1"
)

var (
	uuidCode              *uuid_code.UuidCode = uuid_code.NewDefaultUuidCode()
	ErrRecordAlreadyExist                     = errors.New("exist")
	ErrInvalid                                = errors.New("invalid")
)

type DefaultModel struct {
	ID         uint      `gorm:"primaryKey"`
	CreateDate time.Time `gorm:"autoCreateTime:nano"`
	WriteDate  time.Time `gorm:"autoUpdateTime:nano"`
}
