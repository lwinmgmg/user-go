package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) GetProfile(ctx *gin.Context) {
	sub := getSubject(ctx)
	if resp, err := apiCtrl.Controller.GetProfile(sub.UserID); err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to get profile (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
