package controller

import (
	"errors"

	"github.com/lwinmgmg/user-go/internal/models"
	"gorm.io/gorm"
)

var ErrUserExist = errors.New("user_exist")

type UserSignUpData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserName  string `json:"username"`
	Password  string `json:"password"`
}

func (ctrl *Controller) Signup(userData *UserSignUpData) error {
	user := models.User{
		Username: userData.UserName,
	}
	if err := user.GetUserByUsername(userData.UserName, ctrl.RoDb); err != gorm.ErrRecordNotFound {
		if err == nil {
			return ErrUserExist
		}
		return err
	}
	partner := models.Partner{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Phone:     userData.Phone,
	}
	if err := partner.CheckEmail(ctrl.RoDb); err != nil {
		return err
	}
	if err := partner.CheckPhone(ctrl.RoDb); err != nil {
		return err
	}
	return ctrl.Db.Transaction(func(tx *gorm.DB) error {
		if err := partner.Create(tx); err != nil {
			return err
		}
		user.PartnerID = partner.ID
		return user.Create(tx)
	})
}
