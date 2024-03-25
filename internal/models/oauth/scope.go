package oauth

import (
	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

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

func GetScopesByClientTID(clientTId uint, tx *gorm.DB) ([]Scope, error) {
	cs := []ClientScope{}
	if err := tx.Model(&ClientScope{}).Find(&cs, "client_id=?", clientTId).Error; err != nil {
		return nil, err
	}
	lenCs := len(cs)
	if lenCs == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	sIds := make([]uint, 0, lenCs)
	for _, v := range cs {
		sIds = append(sIds, v.ScopeID)
	}
	scopes := make([]Scope, 0, lenCs)
	err := tx.Model(&Scope{}).Find(&scopes, sIds).Error
	return scopes, err
}
