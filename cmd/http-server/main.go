package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	app := gin.New()
	app.Use(middlewares.LoggerMiddleware, gin.CustomRecovery(middlewares.RecoveryMiddleware))
	app.GET("/", func(ctx *gin.Context) {
		panic(middlewares.NewPanic(http.StatusMethodNotAllowed, 1, "ABC", errors.New("abc")))
	})
	app.Run(fmt.Sprintf("%v:%v", settings.HttpServer.Host, settings.HttpServer.Port))
}
