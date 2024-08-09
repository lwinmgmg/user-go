package migrateme_test

import (
	"errors"
	"testing"

	migrateme "github.com/lwinmgmg/user-go/cmd/migrate-database/migrate-me"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/test"
	"gorm.io/gorm"
)

func TestFindAndSaveIfNotExist(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		partner := models.Partner{
			FirstName: "First",
			LastName:  "Last",
			Email:     "test@test.com",
			Phone:     "093394546",
			Code:      "ncode",
		}
		if err := migrateme.FindAndSaveIfNotExist(&partner, tx); err != nil {
			t.Errorf("Error on creating partner : %v", err)
			return err
		}
		user := models.User{
			Code:      "no_code",
			Username:  "no_user00000000001",
			Password:  []byte("kdjsjfka"),
			PartnerID: partner.ID,
		}
		if err := migrateme.FindAndSaveIfNotExist(&user, tx); err != nil {
			t.Errorf("Error on creating user : %v", err)
			return err
		}
		if err := migrateme.FindAndSaveIfNotExist(&user, tx); err != nil {
			t.Errorf("Error on creating user : %v", err)
			return err
		}
		tx.Rollback()
		tx.AddError(gorm.ErrInvalidTransaction)
		if err := migrateme.FindAndSaveIfNotExist(&user, tx); !errors.Is(err, gorm.ErrInvalidTransaction) {
			t.Errorf("Expected %v : getting %v", gorm.ErrInvalidTransaction, err)
			return err
		}
		return errors.New("to_roll_back")
	})
}
