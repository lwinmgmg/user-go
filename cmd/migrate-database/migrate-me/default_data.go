package migrateme

import (
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"gorm.io/gorm"
)

func MigrateDefaultData(settings *env.Settings, tx *gorm.DB) error {
	user, err := createAdminUser(tx)
	if err != nil {
		return err
	}
	client, err := createDefaultClient(user, settings, tx)
	if err != nil {
		return err
	}
	scopes := []oauth.Scope{
		{
			Name:        "UserRead",
			Description: "Read User Info(User Code, Email, First Name and Last Name)",
			Level:       oauth.SL1,
		},
		{
			Name:        "UserAdmin",
			Description: "Read User Info(User Code, Email, First Name and Last Name)",
			Level:       oauth.SL1,
		},
	}
	for _, scope := range scopes {
		if err := FindAndSaveIfNotExist(&scope, tx); err != nil {
			return err
		}
		if err := FindAndSaveIfNotExist(&oauth.ClientScope{
			ClientID: client.ID,
			ScopeID:  scope.ID,
		}, tx); err != nil {
			return err
		}
	}
	return nil
}
