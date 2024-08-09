package services_test

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/models/oauth"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/hashing"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/test"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func createTestClientUser(tx *gorm.DB) (*oauth.Client, *models.User, error) {
	uuid := hashing.NewUuid4()
	secret := hashing.NewUuid4()
	user, err := models.CreateTestUser("testing", "testing", tx)
	if err != nil {
		return nil, nil, err
	}
	client := oauth.Client{
		Name:          "testing",
		ClientID:      uuid,
		Secret:        secret,
		UserID:        user.ID,
		RedirectUrl:   "http://localhost",
		VerifiedLevel: oauth.SL1,
	}
	if err := tx.Create(&client).Error; err != nil {
		return nil, nil, err
	}
	return &client, user, err
}

func TestGenerateUserJwt(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		// Getting jwt ctrl
		jwtCtrl := jwtctrl.NewJwtCtrl("Testing")
		// Getting client and user
		_, user, err := createTestClientUser(tx)
		if err != nil {
			t.Errorf("Error on creating client, user , : %v", err)
			return err
		}
		// Get formatted key
		formattedKey := services.FormatJwtKey(user.Username, user.Code, string(user.Password), settings.JwtService.Key)
		// Generate token
		jwtStr, err := services.GenerateUserLoginJwt(user.Code, formattedKey, settings, jwtCtrl)
		if err != nil {
			t.Errorf("Error on generating jwt str : %v", err)
			return err
		}
		if jwtStr == "" {
			t.Errorf("Get jwt empty string")
		}
		// Validate with invalid signature
		if _, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(""), nil
		}); !errors.Is(err, jwt.ErrSignatureInvalid) {
			t.Errorf("Expected %v, getting %v", jwt.ErrSignatureInvalid, err)
		}
		// Validate with valid signature
		claim, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), nil
		})
		if err != nil {
			t.Errorf("Getting error on validating jwtToken : %v", err)
			return err
		}
		subStr, err := claim.GetSubject()
		if err != nil {
			t.Errorf("Error on getting subject : %v", err)
			return err
		}
		sub := &jwtctrl.ThirdPartySubject{}
		if err := json.Unmarshal([]byte(subStr), sub); err != nil {
			t.Errorf("Error on unmershaling subject : %v", err)
			return err
		}
		if sub.UserID != user.Code {
			assert.Equal(t, user.Code, sub.UserID, "User ID missmatch")
		}
		// Validate with known error on getKey function
		knownErr := errors.New("known_error")
		if _, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), knownErr
		}); !errors.Is(err, knownErr) {
			t.Errorf("Expected Known Error : getting %v", err)
		}
		// Validate with timeout token
		settings.JwtService.LoginDuration = 0
		jwtStr, err = services.GenerateUserLoginJwt(user.Code, formattedKey, settings, jwtCtrl)
		if err != nil {
			t.Errorf("Error on generating 0 duration jwt str : %v", err)
			return err
		}
		if jwtStr == "" {
			t.Error("Getting empty jwtStr")
		}
		time.Sleep(time.Second)
		if _, err = jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), nil
		}); !errors.Is(err, jwt.ErrTokenExpired) {
			t.Errorf("Expected token expired error, getting %v", err)
		}
		return errors.New("to_roll_back")
	})
}

func TestGenerateThirdpartyJwt(t *testing.T) {
	settings := test.GetTestEnv()
	db, err := services.GetPsql(settings.Db)
	if err != nil {
		t.Error(err.Error())
	}
	db.Transaction(func(tx *gorm.DB) error {
		// Getting jwt ctrl
		jwtCtrl := jwtctrl.NewJwtCtrl("Testing")
		// Getting client and user
		client, user, err := createTestClientUser(tx)
		if err != nil {
			t.Errorf("Error on creating client, user , : %v", err)
			return err
		}
		// Get formatted key
		formattedKey := services.FormatJwtKey(client.ClientID, user.Code, client.Secret, settings.JwtService.Key)
		// Generate token
		jwtStr, err := services.GenerateThirdpartyJwt(user.Code, client.ClientID, formattedKey, settings, jwtCtrl, "UserRead")
		if err != nil {
			t.Errorf("Error on generating jwt str : %v", err)
			return err
		}
		if jwtStr == "" {
			t.Errorf("Get jwt empty string")
		}
		// Validate with invalid signature
		if _, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(""), nil
		}); !errors.Is(err, jwt.ErrSignatureInvalid) {
			t.Errorf("Expected %v, getting %v", jwt.ErrSignatureInvalid, err)
		}
		// Validate with valid signature
		claim, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), nil
		})
		if err != nil {
			t.Errorf("Getting error on validating jwtToken : %v", err)
			return err
		}
		subStr, err := claim.GetSubject()
		if err != nil {
			t.Errorf("Error on getting subject : %v", err)
			return err
		}
		sub := &jwtctrl.ThirdPartySubject{}
		if err := json.Unmarshal([]byte(subStr), sub); err != nil {
			t.Errorf("Error on unmershaling subject : %v", err)
			return err
		}
		if sub.ClientID != client.ClientID {
			assert.Equal(t, sub.ClientID, client.ClientID, "Client ID missmatch")
		}
		if sub.UserID != user.Code {
			assert.Equal(t, user.Code, sub.UserID, "User ID missmatch")
		}
		if sub.Scopes[0] != "UserRead" {
			assert.Equal(t, "UserRead", sub.Scopes[0], "Scope missmatch")
		}
		// Validate with known error on getKey function
		knownErr := errors.New("known_error")
		if _, err := jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), knownErr
		}); !errors.Is(err, knownErr) {
			t.Errorf("Expected Known Error : getting %v", err)
		}
		// Validate with timeout token
		settings.JwtService.LoginDuration = 0
		jwtStr, err = services.GenerateThirdpartyJwt(user.Code, client.ClientID, formattedKey, settings, jwtCtrl, "UserRead")
		if err != nil {
			t.Errorf("Error on generating 0 duration jwt str : %v", err)
			return err
		}
		if jwtStr == "" {
			t.Error("Getting empty jwtStr")
		}
		time.Sleep(time.Second)
		if _, err = jwtCtrl.Validate(jwtStr, func(c jwt.Claims, t *jwt.Token) (any, error) {
			return []byte(formattedKey), nil
		}); !errors.Is(err, jwt.ErrTokenExpired) {
			t.Errorf("Expected token expired error, getting %v", err)
		}
		return errors.New("to_roll_back")
	})
}
