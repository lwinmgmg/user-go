package oauth

import "github.com/lwinmgmg/user-go/internal/models"

type ClientScope struct {
	ClientID uint `gorm:"index"`
	Client   Client
	ScopeID  uint `gorm:"index"`
	Scope    Scope
}

func (cs *ClientScope) TableName() string {
	return models.ComputeTableName("client_scope")
}
