package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (apiCtrl *ApiCtrl) GetProfile(ctx *gin.Context) {
	sub := getSubject(ctx)
	ctx.JSON(http.StatusOK, sub)
}
