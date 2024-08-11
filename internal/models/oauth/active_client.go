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
	activeClient := &ActiveClient{}
	err := tx.Model(activeClient).First(activeClient, "user_id=? AND client_id=?", userId, clientId).Error
	activeClient.RefreshToken = refreshToken
	if errors.Is(err, gorm.ErrRecordNotFound) {
		activeClient.ClientID = clientId
		activeClient.UserID = userId
		err = tx.Create(activeClient).Error
	} else if err == nil {
		err = tx.Updates(activeClient).Error
	}
	return activeClient, err
}

func GetActiveClient(userId, clientId uint, tx *gorm.DB) (*ActiveClient, error) {
	ac := &ActiveClient{}
	if err := tx.Model(ac).First(ac, "user_id=? AND client_id=?", userId, clientId).Error; err != nil {
		return ac, err
	}
	return ac, nil
}
