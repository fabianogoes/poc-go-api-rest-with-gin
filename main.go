package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-demo/go-demo-api/adapter/input/routes"
)

func main() {
	route := gin.Default()

	routes.InitRoutes(route)

	route.Run(":3000")
}
