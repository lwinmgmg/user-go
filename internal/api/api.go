package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
)

type ApiCtrl struct {
	controller.Controller
}

func (apiCtrl *ApiCtrl) RegisterRoutes(router gin.IRoutes) {
	router.POST("/login", apiCtrl.Login)
	router.POST("/signup", apiCtrl.Signup)
}

func NewApiCtrl(ctrl controller.Controller) *ApiCtrl {
	return &ApiCtrl{
		Controller: ctrl,
	}
}
