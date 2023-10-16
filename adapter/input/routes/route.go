package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-demo/go-demo-api/adapter/input/controller"
	"github.com/go-demo/go-demo-api/adapter/output/repository"
	"github.com/go-demo/go-demo-api/application/service"
)

func InitRoutes(route *gin.Engine) {
	customerRepository := repository.NewCustomerInMemoryRepository()
	customerService := service.NewCustomerService(customerRepository)
	customerController := controller.NewCustomerController(customerService)

	route.GET("/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Welcome"}) })
	route.POST("customers", customerController.Create)
	route.GET("/customers", customerController.List)
	route.GET("/customers/:id", customerController.FindById)
	route.PUT("/customers/:id", customerController.Update)
	route.DELETE("/customers/:id", customerController.Delete)
}
