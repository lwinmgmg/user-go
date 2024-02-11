package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) ResendOtp(ctx *gin.Context) {
	data := &controller.ResendOtpRequest{}
	if err := ctx.ShouldBind(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	if loginTkn, err := apiCtrl.Controller.ResendOtp(data); err != nil {
		switch err {
		case controller.ErrOtpNotFound:
			panic(middlewares.NewPanic(http.StatusNotFound, 1, "User Not Found", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to signup (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, loginTkn)
	}
}
