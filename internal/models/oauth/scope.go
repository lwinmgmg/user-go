package oauth

import "github.com/lwinmgmg/user-go/internal/models"

type Scope struct {
	models.DefaultModel
	Name        string
	Description string
}

func (cs *Scope) TableName() string {
	return models.ComputeTableName("scope")
}
