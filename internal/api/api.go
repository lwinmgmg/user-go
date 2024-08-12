package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

type ApiCtrl struct {
	controller.Controller
}

func (apiCtrl *ApiCtrl) RegisterRoutes(router gin.IRouter) {
	funcRouter := router.Group("/api/v1/func/user")
	funcRouter.POST("/login", apiCtrl.Login)
	funcRouter.POST("/signup", apiCtrl.Signup)
	funcRouter.POST("/otp_auth", apiCtrl.OtpAuth)
	funcRouter.POST("/resend_otp", apiCtrl.ResendOtp)

	userRouter := router.Group("/api/v1/user")
	userRouter.GET("/profile", apiCtrl.AuthMiddleware, apiCtrl.GetProfile)
	userRouter.GET("/profile_detail", apiCtrl.AuthMiddleware, apiCtrl.GetProfileDetail)
	userRouter.GET("/email_confirm", apiCtrl.AuthMiddleware, apiCtrl.EmailConfirm)
	userRouter.GET("/phone_confirm", apiCtrl.AuthMiddleware, apiCtrl.PhoneConfirm)
	userRouter.GET("/enable_2fa", apiCtrl.AuthMiddleware, apiCtrl.Enable2FA)
	userRouter.GET("/enable_authenticator", apiCtrl.AuthMiddleware, apiCtrl.EnableAuthenticator)
	userRouter.POST("/oauth/thirdparty/generate", apiCtrl.AuthMiddleware, apiCtrl.GenerateThirdPartyToken)
	userRouter.POST("/oauth/thirdparty/check", apiCtrl.CheckThirdPartyTkn)
	userRouter.POST("/oauth/thirdparty/get", apiCtrl.GetThirdPartyToken)
}

func NewApiCtrl(ctrl controller.Controller) *ApiCtrl {
	return &ApiCtrl{
		Controller: ctrl,
	}
}

func SetupRouter(settings *env.Settings) *gin.Engine {
	app := gin.New()
	corsConf := cors.Config{
		AllowAllOrigins: true,
		AllowHeaders: []string{
			"content-type",
		},
	}
	app.Use(middlewares.LoggerMiddleware, gin.CustomRecovery(middlewares.RecoveryMiddleware), cors.New(corsConf))
	apiCtrl := NewApiCtrl(*controller.NewContoller(settings))
	apiCtrl.RegisterRoutes(app.Group("/user"))
	return app
}
