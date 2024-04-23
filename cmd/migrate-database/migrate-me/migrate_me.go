package migrateme

import (
	"errors"

	"gorm.io/gorm"
)

func FindAndSaveIfNotExist(obj any, tx *gorm.DB) error {
	if err := tx.Where(obj).First(obj).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return tx.Create(obj).Error
		}
		return err
	}
	return nil
}
