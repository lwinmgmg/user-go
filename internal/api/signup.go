package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) Signup(ctx *gin.Context) {
	data := &controller.UserSignUpData{}
	if err := ctx.ShouldBindJSON(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	if err := data.Validate(); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 2, "Invalid data format", err))
	}
	if err := apiCtrl.Controller.Signup(data); err != nil {
		switch err {
		case controller.ErrUserExist:
			panic(middlewares.NewPanic(http.StatusBadRequest, 4, "User already exist", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to signup", err))
	}
	ctx.JSON(http.StatusOK, middlewares.DefResp{
		Code:    1,
		Message: "Successfully Signup",
	})
}
