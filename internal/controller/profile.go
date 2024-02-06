package controller

import "github.com/lwinmgmg/user-go/internal/models"

type Profile struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	ImageUrl  string `json:"image_url"`
	UserCode  string `json:"user_id"`
}

func (ctrl *Controller) GetProfile(userCode string) (*Profile, error) {
	user := models.User{}
	if _, err := user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return nil, err
	}
	return &Profile{
		FirstName: user.Partner.FirstName,
		LastName:  user.Partner.LastName,
		Email:     user.Partner.Email,
		UserCode:  userCode,
	}, nil
}
