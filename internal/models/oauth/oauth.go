package oauth

import "gorm.io/gorm"

func InitDb(db *gorm.DB) error {
	client := &Client{}
	scope := &Scope{}
	cs := &ClientScope{}
	if err := db.Migrator().AutoMigrate(
		client,
		scope,
		cs,
	); err != nil {
		return err
	}
	return nil
}
