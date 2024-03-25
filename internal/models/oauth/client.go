package oauth

import (
	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

type Client struct {
	models.DefaultModel
	Name          string
	ClientID      string `gorm:"index"`
	Secret        string
	UserID        uint `gorm:"index"`
	User          models.User
	Verified      bool       `gorm:"default: false;"`
	VerifiedLevel ScopeLevel `gorm:"default: '1'"`
	RedirectUrl   string
}

func (cs *Client) TableName() string {
	return models.ComputeTableName("client")
}

func GetClientByCId(cid string, tx *gorm.DB) (*Client, error) {
	client := &Client{}
	if err := tx.Model(client).First(client, "client_id=?", cid).Error; err != nil {
		return nil, err
	}
	return client, nil
}
