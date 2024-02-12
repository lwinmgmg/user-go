package controller

import "github.com/lwinmgmg/user-go/internal/models"

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
	
	return
}
