package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/controller"
	jwtctrl "github.com/lwinmgmg/user-go/pkg/jwt-ctrl"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

type ThirdpartyTokenRequest struct {
	Code        string   `form:"code"`
	ClientID    string   `form:"cid" binding:"required,min=3"`
	Scopes      []string `form:"scp"`
	RedirectUrl string   `form:"rurl" binding:"required,min=3"`
}

func (apiCtrl *ApiCtrl) GenerateThirdPartyToken(ctx *gin.Context) {
	sub := getUserSubject[jwtctrl.UserSubject](ctx)
	data := &ThirdpartyTokenRequest{}
	if err := ctx.ShouldBind(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	if resp, err := apiCtrl.Controller.GenerateThirdPartyToken(sub.UserID, data.ClientID, data.RedirectUrl, data.Scopes...); err != nil {
		if errors.Is(err, controller.ErrNoScope) {
			panic(middlewares.NewPanic(http.StatusNotFound, 0, "No Scope found", err))
		}
		if errors.Is(err, controller.ErrUnauthorizedScope) {
			panic(middlewares.NewPanic(http.StatusUnauthorized, 0, "No Authorized Scopes", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to generate thirdparty token", err))
	} else {
		ctx.JSON(http.StatusOK, resp)
	}
}
