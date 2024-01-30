package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/internal/api"
	"github.com/lwinmgmg/user-go/internal/controller"
	"github.com/lwinmgmg/user-go/pkg/middlewares"
)

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}
	app := gin.New()
	app.Use(middlewares.LoggerMiddleware, gin.CustomRecovery(middlewares.RecoveryMiddleware))
	apiCtrl := api.NewApiCtrl(*controller.NewContoller(settings))
	apiCtrl.RegisterRoutes(app.Group("/user"))
	app.Run(fmt.Sprintf("%v:%v", settings.HttpServer.Host, settings.HttpServer.Port))
}
