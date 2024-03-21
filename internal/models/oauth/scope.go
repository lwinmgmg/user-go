package oauth

import "github.com/lwinmgmg/user-go/internal/models"

type ScopeLevel string

const (
	SL1 ScopeLevel = "1"
	SL2 ScopeLevel = "2"
	SL3 ScopeLevel = "3"
)

type Scope struct {
	models.DefaultModel
	Name        string
	Description string
	Level       ScopeLevel
}

func (cs *Scope) TableName() string {
	return models.ComputeTableName("scope")
}
