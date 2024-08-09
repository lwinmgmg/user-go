package oauth

import (
	"errors"

	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

type ActiveClient struct {
	models.DefaultModel
	ClientID     uint
	Client       Client
	UserID       uint
	User         models.User
	RefreshToken string
}

func (ac *ActiveClient) TableName() string {
	return models.ComputeTableName("active_client")
}

func GetActiveClientCreateIfNotExist(userId, clientId uint, refreshToken string, tx *gorm.DB) (*ActiveClient, error) {
	acs := &ActiveClient{
		UserID:       userId,
		ClientID:     clientId,
		RefreshToken: refreshToken,
	}
	if err := tx.Model(&ActiveClient{}).First(acs, "user_id=? AND client_id=?", userId, clientId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if err := tx.Create(acs).Error; err != nil {
			return acs, err
		}
	}
	return acs, nil
}

func GetActiveClient(userId, clientId uint, tx *gorm.DB) (*ActiveClient, error) {
	ac := &ActiveClient{}
	if err := tx.Model(ac).First(ac, "user_id=? AND client_id=?", userId, clientId).Error; err != nil {
		return ac, err
	}
	return ac, nil
}
