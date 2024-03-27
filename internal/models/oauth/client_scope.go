package oauth

import (
	"errors"
	"slices"

	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

var (
	ErrScopeIsNotValid = errors.New("scope_not_valid")
)

type ClientScope struct {
	ClientID uint `gorm:"index"`
	Client   Client
	ScopeID  uint `gorm:"index"`
	Scope    Scope
}

func (cs *ClientScope) TableName() string {
	return models.ComputeTableName("client_scope")
}

func CheckClientScope(clientId uint, tx *gorm.DB, scopeNames ...string) ([]Scope, error) {
	scopes := make([]Scope, 0, len(scopeNames))
	if err := tx.Model(&Scope{}).Find(&scopes, "name in ?", scopeNames).Error; err != nil {
		return scopes, err
	}
	validScopeIds := []uint{}
	if err := tx.Model(&ClientScope{}).Select("scope_id").Find(&validScopeIds, "client_id=?", clientId).Error; err != nil {
		return scopes, err
	}
	for _, scope := range scopes {
		if !slices.Contains(validScopeIds, scope.ID) {
			return scopes, ErrScopeIsNotValid
		}
	}
	return scopes, nil
}
