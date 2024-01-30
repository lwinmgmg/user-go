package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
	"gorm.io/gorm"
)

func (apiCtrl *ApiCtrl) Signup(ctx *gin.Context) {
	if err := apiCtrl.Controller.Signup(&controller.UserSignUpData{}); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			panic(middlewares.NewPanic(http.StatusBadRequest, 1, "Failed to signup", err))
		case controller.ErrUserExist:
			panic(middlewares.NewPanic(http.StatusBadRequest, 1, "Failed to signup", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 1, "Failed to signup", err))
	}
	ctx.JSON(http.StatusOK, middlewares.DefResp{
		Code:    1,
		Message: "Successfully Signup",
	})
}
