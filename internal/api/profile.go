package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) GetProfile(ctx *gin.Context) {
	sub := getUserSubject[jwtctrl.UserSubject](ctx)
	if resp, err := apiCtrl.Controller.GetProfile(sub.UserID); err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to get profile (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
