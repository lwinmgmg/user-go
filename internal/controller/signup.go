package controller

import (
	"errors"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

var ErrUserExist = errors.New("user_exist")

type UserSignUpData struct {
	FirstName string `json:"first_name" binding:"required,min=3"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone" binding:"required"`
	UserName  string `json:"username" binding:"required,min=3"`
	Password  string `json:"password"`
}

func (data *UserSignUpData) Validate() error {
	return nil
}

func (ctrl *Controller) Signup(userData *UserSignUpData) (loginTkn LoginToken, err error) {
	hashPass, err := hashing.Hash256(userData.Password)
	if err != nil {
		return
	}
	user := models.User{
		Username: userData.UserName,
		Password: hashPass,
	}
	if err = user.GetUserByUsername(userData.UserName, ctrl.RoDb); err != gorm.ErrRecordNotFound {
		if err == nil {
			return loginTkn, ErrUserExist
		}
		return
	}
	partner := models.Partner{
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
		Phone:     userData.Phone,
	}
	if err = partner.CheckEmail(ctrl.RoDb); err != nil {
		return
	}
	if err = partner.CheckPhone(ctrl.RoDb); err != nil {
		return
	}
	if err = ctrl.Db.Transaction(func(tx *gorm.DB) error {
		if err := partner.Create(tx); err != nil {
			return err
		}
		user.PartnerID = partner.ID
		return user.Create(tx)
	}); err != nil {
		return
	}
	formattedKey := services.FormatJwtKey(user.Username, user.Code, string(user.Password), ctrl.Setting.JwtService.Key)
	jwtToken, err := services.GenerateUserLoginJwt(user.Code, formattedKey, ctrl.Setting, ctrl.JwtCtrl)
	if err != nil {
		return
	}
	loginTkn.Value = jwtToken
	return
}
