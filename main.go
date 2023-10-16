package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-demo/go-demo-api/adapter/input/routes"
	"github.com/go-demo/go-demo-api/configuration/logger"
)

func main() {
	logger.Info("Application running...")
	route := gin.Default()

	routes.InitRoutes(route)

	route.Run(":3000")
}
