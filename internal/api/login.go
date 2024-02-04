package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/internal/models"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
	"gorm.io/gorm"
)

type LoginData struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=3"`
}

func (apiCtrl *ApiCtrl) Login(ctx *gin.Context) {
	data := &LoginData{}
	if err := ctx.ShouldBindJSON(data); err != nil {
		panic(middlewares.NewPanic(http.StatusUnprocessableEntity, 1, "Wrong data format", err))
	}
	if loginTkn, err := apiCtrl.Controller.Login(data.Username, data.Password, &models.User{}); err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			panic(middlewares.NewPanic(http.StatusNotFound, 1, "User Not Found", err))
		case models.ErrWrongPassword:
			panic(middlewares.NewPanic(http.StatusUnauthorized, 2, "Wrong Password", err))
		}
		panic(middlewares.NewPanic(http.StatusBadRequest, 0, "Failed to signup (Unknown)", err))
	} else {
		ctx.JSON(http.StatusOK, loginTkn)
	}
}
