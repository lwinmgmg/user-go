package controller

import "github.com/lwinmgmg/user-go/internal/models"

type ProfileDetail struct {
	FirstName        string `json:"firstname"`
	LastName         string `json:"lastname"`
	Email            string `json:"email"`
	Phone            string `json:"phone"`
	ImageUrl         string `json:"image_url"`
	UserCode         string `json:"user_id"`
	IsEmailConfirmed bool   `json:"is_email"`
	IsPhoneConfirmed bool   `json:"is_phone"`
	IsAuthenticator  bool   `json:"is_auth"`
	Is2Fa            bool   `json:"is_2fa"`
}

func (ctrl *Controller) GetProfileDetail(userCode string) (*ProfileDetail, error) {
	user := models.User{}
	if _, err := user.GetPartnerByCode(userCode, ctrl.RoDb); err != nil {
		return nil, err
	}
	return &ProfileDetail{
		FirstName:        user.Partner.FirstName,
		LastName:         user.Partner.LastName,
		Email:            user.Partner.Email,
		Phone:            user.Partner.Phone,
		UserCode:         userCode,
		IsEmailConfirmed: user.Partner.IsEmailConfirmed,
		IsPhoneConfirmed: user.Partner.IsPhoneConfirmed,
		IsAuthenticator:  user.IsAuthenticator,
		Is2Fa:            user.OtpUrl != "",
	}, nil
}
