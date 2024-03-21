package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) PhoneConfirm(ctx *gin.Context) {
	sub := getUserSubject[jwtctrl.UserSubject](ctx)
	if resp, err := apiCtrl.Controller.PhoneConfirm(sub.UserID); err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to confirm (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
