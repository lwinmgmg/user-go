package models

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/lwinmgmg/user-go/pkg/hashing"
	"gorm.io/gorm"
)

const (
	USER_CODE_LENGTH int = 10
)

var (
	ErrWrongPassword = errors.New("wrong_password")
)

type User struct {
	DefaultModel
	Code            string  `gorm:"uniqueIndex; index; not null; size:10;"`
	Username        string  `gorm:"uniqueIndex; index; not null; size:32;"`
	Password        []byte  `gorm:"size:256;"`
	PartnerID       uint    `gorm:"uniqueIndex; not null; index"`
	Partner         Partner `gorm:"foreignKey:PartnerID"`
	OtpUrl          string  `gorm:"size:256; index;"`
	IsAuthenticator bool    `gorm:"default:false;"`
}

func (user *User) TableName() string {
	return computeTableName("user")
}

func (user *User) GetSequence() string {
	return "user_sequence"
}

func (user *User) NextCode(db *gorm.DB) string {
	var nextSequence int
	db.Raw(fmt.Sprintf("SELECT nextval('%v');", user.GetSequence())).Scan(&nextSequence)
	return uuidCode.ConvertCode(nextSequence, USER_CODE_LENGTH)
}

func (user *User) Create(tx *gorm.DB) error {
	if strings.TrimSpace(user.Username) == "" {
		return ErrInvalid
	}
	user.Code = user.NextCode(tx)
	return tx.Create(user).Error
}

func (user *User) Authenticate(tx *gorm.DB, username, password string) error {
	if err := tx.Where(&User{
		Username: username,
	}).First(user).Error; err != nil {
		return err
	}
	if user.Username == "" {
		return errors.New("user not found")
	}
	hPass, err := hashing.Hash256(password)
	if err != nil {
		return err
	}
	if !bytes.Equal(user.Password, hPass) {
		return ErrWrongPassword
	}
	return nil
}

func (user *User) GetPartnerByUsername(username string, tx *gorm.DB) (*Partner, error) {
	if err := tx.Where(&User{
		Username: username,
	}).First(user).Error; err != nil {
		return &user.Partner, err
	}
	if err := tx.First(&user.Partner, user.PartnerID).Error; err != nil {
		return &user.Partner, err
	}
	return &user.Partner, nil
}

func (user *User) GetUserByUsername(username string, tx *gorm.DB) error {
	return tx.Where(&User{
		Username: username,
	}).First(user).Error
}

func (user *User) GetUserByCode(code string, tx *gorm.DB) error {
	return tx.Where(&User{
		Code: code,
	}).First((user)).Error
}

func (user *User) GetPartnerByCode(code string, tx *gorm.DB) (*Partner, error) {
	if err := tx.Where(&User{
		Code: code,
	}).First(user).Error; err != nil {
		return &user.Partner, err
	}
	if err := tx.First(&user.Partner, user.PartnerID).Error; err != nil {
		return &user.Partner, err
	}
	return &user.Partner, nil
}

func (user *User) Login() error {
	return nil
}

func (user *User) SetOtpUrl(url string, tx *gorm.DB) error {
	user.OtpUrl = url
	return tx.Save(user).Error
}

func (user *User) ChangePassword(newPassword string) error {
	return nil
}

func (user *User) ChangeEmail(newEmail string) error {
	return nil
}

func (user *User) ChangePhone(newPhone string) error {
	return nil
}

func (user *User) SetIsAuthenticator(input bool, tx *gorm.DB) error {
	user.IsAuthenticator = input
	return tx.Save(user).Error
}

func GetPasswordByUserCode(code string, tx *gorm.DB) (string, error) {
	var password string
	if err := tx.Model(&User{}).Select("password").Where("code = ?", code).First(&password).Error; err != nil {
		return "", err
	}
	return password, nil
}

func CreateTestUser(username, password string, tx *gorm.DB) (*User, error) {
	partner := Partner{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "test@mail.com",
	}
	if err := partner.Create(tx); err != nil {
		return nil, err
	}
	hPass, err := hashing.Hash256(password)
	if err != nil {
		return nil, err
	}
	user := User{
		Username:  username,
		Password:  hPass,
		PartnerID: partner.ID,
	}
	if err := user.Create(tx); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUsers(tx *gorm.DB) (users []User, err error) {
	err = tx.Model(&User{}).Find(&users).Error
	return
}

func GetUserById(id uint, tx *gorm.DB) (user *User, err error) {
	user = &User{}
	err = tx.Model(user).First(user, id).Error
	return
}
