package main

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/lwinmgmg/user-go/env"
	"github.com/lwinmgmg/user-go/graph"
	"github.com/lwinmgmg/user-go/internal/api"
	"github.com/lwinmgmg/user-go/internal/controller"
)

func graphqlHandler(ctrl *controller.Controller) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		RoDb: ctrl.RoDb,
		Db:   ctrl.Db,
	}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler(name, path string) gin.HandlerFunc {
	h := playground.Handler(name, path)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	settings, err := env.LoadSettings()
	if err != nil {
		panic(err)
	}

	// Graphql
	// app.Handle("GET", "/graphql", playgroundHandler("Graph QL UI", "/graphql"))
	// app.Handle("POST", "/graphql", graphqlHandler(&apiCtrl.Controller))
	app := api.SetupRouter(settings)
	app.Run(fmt.Sprintf("%v:%v", settings.HttpServer.Host, settings.HttpServer.Port))
}
