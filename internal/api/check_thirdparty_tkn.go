package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

type TknCheckReq struct {
	AccessToken string `form:"tkn" binding:"required"`
}

func (apiCtrl *ApiCtrl) CheckThirdPartyTkn(ctx *gin.Context) {
	data := TknCheckReq{}
	if err := ctx.ShouldBind(&data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	sub, err := apiCtrl.Controller.CheckThirdPartyTkn(data.AccessToken)
	if err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to signup (Unknown)", err))
	}
	ctx.JSON(http.StatusOK, &sub)
}
