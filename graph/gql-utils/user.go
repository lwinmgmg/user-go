package gqlutils

import (
	"github.com/lwinmgmg/user-go/graph/model"
	"github.com/lwinmgmg/user-go/internal/models"
)

func ContructUser(user *models.User) *model.User {
	return &model.User{
		ID:              int(user.ID),
		Username:        user.Username,
		Password:        string(user.Password),
		Code:            user.Code,
		PartnerID:       int(user.PartnerID),
		OtpURL:          &user.OtpUrl,
		IsAuthenticator: user.IsAuthenticator,
	}
}
