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

func (apiCtrl *ApiCtrl) RegisterRoutes(router gin.IRoutes) {
	router.POST("/login", apiCtrl.Login)
	router.POST("/signup", apiCtrl.Signup)
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
