package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-demo/go-demo-api/application/entity"
	"github.com/go-demo/go-demo-api/application/service"
	"github.com/go-demo/go-demo-api/configuration/logger"
)

type CustomerResponse struct {
	ID   string
	Name string
}

type CustomerCreateRequest struct {
	Name string `json:"name"`
}

type customerController struct {
	CustomerService *service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) *customerController {
	return &customerController{CustomerService: customerService}
}

func (cc *customerController) List(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cc.CustomerService.List())
}

func (cc *customerController) Create(c *gin.Context) {
	var request CustomerCreateRequest
	var err error

	if err = c.BindJSON(&request); err != nil {
		logger.Error("Error on bind request", err)
		return
	}

	if err = cc.CustomerService.Create(request.Name); err != nil {
		logger.Error("Error on create customer", err)
		return
	}

	c.IndentedJSON(http.StatusCreated, nil)
}

func (cc *customerController) FindById(c *gin.Context) {
	id := c.Param("id")
	if customer := cc.CustomerService.FindById(id); customer != nil {
		c.IndentedJSON(http.StatusOK, CustomerResponse{ID: customer.ID, Name: customer.Name})
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func (cc *customerController) Update(c *gin.Context) {
	id := c.Param("id")
	var err error

	if cust := cc.CustomerService.FindById(id); cust == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("customer: [%s] not found!", id)})
		return
	}

	var customer entity.Customer
	if err = c.BindJSON(&customer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err = cc.CustomerService.Update(id, &customer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Customer: [%s] updated sucessfully", id)})
}

func (cc *customerController) Delete(c *gin.Context) {
	id := c.Param("id")
	if cost := cc.CustomerService.FindById(id); cost == nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Customer: [%s] not found", id)})
		return
	}

	if err := cc.CustomerService.Delete(id); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("Customer: [%s] deleted sucessfully", id)})
}
