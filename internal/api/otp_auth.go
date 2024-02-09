package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/internal/services"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) OtpAuth(ctx *gin.Context) {
	data := &controller.OtpAuth{}
	if err := ctx.ShouldBind(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	user := models.User{}
	tkn, confirmType, err := apiCtrl.Controller.OtpAuth(
		data, &user)
	if err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Otp authentication failed (Unknown)", err))
	}
	if confirmType == services.OtpLogin {
		ctx.JSON(http.StatusOK, tkn)
		return
	}
	ctx.JSON(http.StatusOK, middlewares.DefResp{
		Code:    1,
		Message: "Successfully authenticated",
	})
}
