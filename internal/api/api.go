package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

type ApiCtrl struct {
	controller.Controller
}

func (apiCtrl *ApiCtrl) Login(ctx *gin.Context) {
	if err := apiCtrl.Controller.Login("", "", &models.User{}); err != nil {
		if err != nil {
			panic(middlewares.NewPanic(http.StatusBadRequest, 1, "Failed to login", err))
		}
	}
	ctx.JSON(http.StatusOK, middlewares.DefResp{
		Code:    1,
		Message: "Successfully Login",
	})
}
