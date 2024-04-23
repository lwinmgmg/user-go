package gqlutils

import (
	"github.com/lwinmgmg/user-go/graph/model"
	"github.com/lwinmgmg/user-go/internal/models"
)

func ContructPartner(ptner *models.Partner) *model.Partner {
	return &model.Partner{
		ID:               int(ptner.ID),
		FirstName:        ptner.FirstName,
		LastName:         &ptner.LastName,
		Email:            ptner.Email,
		Phone:            ptner.Phone,
		IsEmailConfirmed: ptner.IsEmailConfirmed,
		IsPhoneConfirmed: ptner.IsPhoneConfirmed,
		Code:             ptner.Code,
	}
}
