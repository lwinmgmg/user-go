package oauth

import (
	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

type ActiveClientScope struct {
	ActiveClientID uint `gorm:"index;"`
	ActiveClient   ActiveClient
	ScopeID        uint `gorm:"index;"`
	Scope          Scope
}

func (ac *ActiveClientScope) TableName() string {
	return models.ComputeTableName("active_client_scope")
}

func GetAcsByUserClientId(activeClientId uint, tx *gorm.DB) ([]Scope, error) {
	scopeIds := []uint{}
	if err := tx.Model(&ActiveClientScope{}).Select("scope_id").Find(&scopeIds, "active_client_id=?", activeClientId).Error; err != nil {
		return nil, err
	}
	if len(scopeIds) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	acs := make([]Scope, 0, len(scopeIds))
	if err := tx.Model(&Scope{}).Find(&acs, scopeIds).Error; err != nil {
		return nil, err
	}
	return acs, nil
}

// Delete old active client scopes and add the new active client scopes
func CreateActiveClientScope(activeClientId uint, tx *gorm.DB, scopes ...string) error {
	if err := tx.Model(&ActiveClientScope{}).Delete(nil, "active_client_id=?", activeClientId).Error; err != nil {
		return err
	}
	scopeCount := len(scopes)
	scopeList := make([]Scope, 0, scopeCount)
	if err := tx.Model(&Scope{}).Find(&scopeList, "name in ?", scopes).Error; err != nil {
		return err
	}
	acsList := make([]ActiveClientScope, 0, scopeCount)
	for _, v := range scopeList {
		acsList = append(acsList, ActiveClientScope{
			ActiveClientID: activeClientId,
			ScopeID:        v.ID,
		})
	}
	return tx.Model(&ActiveClientScope{}).Create(&acsList).Error
}
