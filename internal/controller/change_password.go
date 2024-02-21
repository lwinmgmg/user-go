package controller

import (
	"errors"
	"slices"

	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

var (
	ErrPasswordNotMatch = errors.New("password_not_match")
)

type PasswordChange struct {
	OldPassword string `form:"old_pass"`
	NewPassword string `form:"new_pass"`
}

func (ctrl *Controller) ChangePassword(userCode string, data *PasswordChange) (logintkn LoginToken, err error) {
	user := models.User{}
	_, err = user.GetPartnerByCode(userCode, ctrl.RoDb)
	if err != nil {
		return
	}
	var oldPwd []byte
	oldPwd, err = hashing.Hash256(data.OldPassword)
	if err != nil {
		return
	}
	if slices.Compare(oldPwd, user.Password) != 0 {
		err = ErrPasswordNotMatch
		return
	}
	err = ctrl.Db.Transaction(
		func(tx *gorm.DB) error {
			return tx.Save(&user).Error
		},
	)
	return
}
