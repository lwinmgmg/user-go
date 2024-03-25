package oauth

import (
	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

type ActiveClient struct {
	ClientID uint
	Client   Client
	UserID   uint
	User     models.User
}

func (ac *ActiveClient) TableName() string {
	return models.ComputeTableName("active_client")
}

func GetActiveClients(userId uint, tx *gorm.DB) ([]ActiveClient, error) {
	acs := []ActiveClient{}
	err := tx.Model(&ActiveClient{}).Find(&acs, "user_id=?", userId).Error
	return acs, err
}
