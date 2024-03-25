package oauth

import "github.com/lwinmgmg/user-go/internal/models"

type ActiveClientScope struct {
	ActiveClientID uint
	ActiveClient   ActiveClient
	ScopeID        uint
	Scope          Scope
}

func (ac *ActiveClientScope) TableName() string {
	return models.ComputeTableName("active_client_scope")
}
