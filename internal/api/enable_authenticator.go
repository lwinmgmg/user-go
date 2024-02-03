package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func (apiCtrl *ApiCtrl) EnableAuthenticator(ctx *gin.Context) {
	sub := getSubject(ctx)
	if resp, err := apiCtrl.Controller.EnableAuthenticator(sub.UserID); err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to enable Authenticator (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
