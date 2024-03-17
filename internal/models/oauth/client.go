package oauth

import "github.com/lwinmgmg/user-go/internal/models"

type Client struct {
	models.DefaultModel
	Name     string
	ClientID string `gorm:"index"`
	Secret   string
	UserID   int `gorm:"index"`
	User     models.User
}

func (cs *Client) TableName() string {
	return models.ComputeTableName("client")
}
