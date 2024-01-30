package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
	"gorm.io/gorm"
)

func (apiCtrl *ApiCtrl) Login(ctx *gin.Context) {
	if err := apiCtrl.Controller.Login("", "", &models.User{}); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			panic(middlewares.NewPanic(http.StatusNotFound, 1, "User Not Found", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to signup (Unknown)", err))
	}
	ctx.JSON(http.StatusOK, middlewares.DefResp{
		Code:    1,
		Message: "Successfully Login",
	})
}
