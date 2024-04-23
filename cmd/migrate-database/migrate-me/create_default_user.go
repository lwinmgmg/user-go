package migrateme

import (
	"errors"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

func createAdminUser(tx *gorm.DB) (*models.User, error) {
	password, err := hashing.Hash256("admin")
	if err != nil {
		return nil, err
	}
	user := models.User{
		Username: "admin",
		Code:     "admin",
		Password: password,
	}
	err = tx.Model(&models.User{}).Where("code=?", "admin").First(&user).Error
	if err == nil {
		return &user, err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	partner := models.Partner{
		FirstName: "Administrator",
		Code:      "admin",
	}
	if err := tx.Create(&partner).Error; err != nil {
		return nil, err
	}
	user.PartnerID = partner.ID
	if err := tx.Create(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}
