package models

import (
	"fmt"

	"gorm.io/gorm"
)

const (
	PARTNER_CODE_LENGTH int = 5
)

type Partner struct {
	DefaultModel
	FirstName        string `gorm:"size:20;"`
	LastName         string `gorm:"size:20;"`
	Email            string `gorm:"index; size:30;"`
	Phone            string `gorm:"index; size:15;"`
	IsEmailConfirmed bool   `gorm:"default:false"`
	IsPhoneConfirmed bool   `gorm:"default:false"`
	Code             string `gorm:"uniqueIndex; not null; size:5;"`
}

func (partner *Partner) TableName() string {
	return computeTableName("partner")
}

func (partner *Partner) GetSequence() string {
	return "partner_sequence"
}

func (partner *Partner) NextCode(db *gorm.DB) string {
	var nextSequence int
	db.Raw(fmt.Sprintf("SELECT nextval('%v');", partner.GetSequence())).Scan(&nextSequence)
	return uuidCode.ConvertCode(nextSequence, PARTNER_CODE_LENGTH)
}

func (partner *Partner) Create(tx *gorm.DB) error {
	partner.Code = partner.NextCode(tx)
	return tx.Create(partner).Error
}

func (partner *Partner) CheckEmail(tx *gorm.DB) error {
	if err := tx.Model(partner).Where(Partner{
		Email:            partner.Email,
		IsEmailConfirmed: true,
	}).First(&Partner{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return ErrRecordAlreadyExist
}

func (partner *Partner) CheckPhone(tx *gorm.DB) error {
	if err := tx.Model(partner).Where(Partner{
		Phone:            partner.Phone,
		IsPhoneConfirmed: true,
	}).First(&Partner{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}
	return ErrRecordAlreadyExist
}

func (partner *Partner) GetPartnerByID(id uint, tx *gorm.DB) error {
	return tx.First(partner, id).Error
}

func (partner *Partner) SetEmailConfirm(input bool, tx *gorm.DB) error {
	partner.IsEmailConfirmed = input
	return tx.Save(partner).Error
}

func (partner *Partner) SetPhoneConfirm(input bool, tx *gorm.DB) error {
	partner.IsPhoneConfirmed = input
	return tx.Save(partner).Error
}

func GetPartners(tx *gorm.DB) (partners []Partner, err error) {
	err = tx.Model(&Partner{}).Find(&partners).Error
	return
}

func GetPartnerByID(id uint, tx *gorm.DB) (partner *Partner, err error) {
	partner = &Partner{}
	err = tx.Model(partner).First(partner, id).Error
	return
}
