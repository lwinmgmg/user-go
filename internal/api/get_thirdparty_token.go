package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

type ThirdPartyTknGet struct {
	Code         string `form:"code"`
	ClientID     string `form:"cid"`
	ClientSecret string `form:"secret"`
	UserCode     string `form:"uid"`
}

func (apiCtrl *ApiCtrl) GetThirdPartyToken(ctx *gin.Context) {
	data := &ThirdPartyTknGet{}
	if err := ctx.ShouldBind(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	if tkn, err := apiCtrl.Controller.GetThirdPartyToken(data.Code, data.UserCode, data.ClientID, data.ClientSecret); err != nil {
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to generate thirdparty token", err))
	} else {
		ctx.JSON(http.StatusOK, tkn)
	}
}
