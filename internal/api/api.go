package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/controller"
)

type ApiCtrl struct {
	controller.Controller
	Settings env.Settings
}

func (apiCtrl *ApiCtrl) RegisterRoutes(router gin.IRouter) {
	funcRouter := router.Group("/api/v1/func/user")
	funcRouter.POST("/login", apiCtrl.Login)
	funcRouter.POST("/signup", apiCtrl.Signup)
	funcRouter.POST("/otp_auth", apiCtrl.OtpAuth)

	userRouter := router.Group("/api/v1/user")
	userRouter.GET("/profile", apiCtrl.AuthMiddleware, apiCtrl.GetProfile)
	userRouter.GET("/profile_detail", apiCtrl.AuthMiddleware, apiCtrl.GetProfileDetail)
	userRouter.GET("/email_confirm", apiCtrl.AuthMiddleware, apiCtrl.EmailConfirm)
	userRouter.GET("/phone_confirm", apiCtrl.AuthMiddleware, apiCtrl.PhoneConfirm)
	userRouter.GET("/enable_2fa", apiCtrl.AuthMiddleware, apiCtrl.Enable2FA)
	userRouter.GET("/enable_authenticator", apiCtrl.AuthMiddleware, apiCtrl.EnableAuthenticator)
}

func NewApiCtrl(ctrl controller.Controller) *ApiCtrl {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	return &ApiCtrl{
		Controller: ctrl,
		Settings:   settings,
	}
}
