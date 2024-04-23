package migrateme

import (
	"errors"

	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

func createDefaultClient(user *models.User, settings *env.Settings, tx *gorm.DB) (*oauth.Client, error) {
	client := oauth.Client{
		Name: settings.Service,
	}
	if err := tx.Where("name = ?", settings.Service).First(&client).Error; err == nil {
		return &client, err
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return &client, err
	}
	cid := hashing.NewUuid4()
	password, err := hashing.Hash256Hex(hashing.NewUuid4())
	if err != nil {
		return nil, err
	}
	client.ClientID = cid
	client.Secret = password
	client.UserID = user.ID
	client.Verified = true
	client.VerifiedLevel = oauth.SL3
	client.RedirectUrl = "http://localhost"
	return &client, tx.Create(&client).Error
}
