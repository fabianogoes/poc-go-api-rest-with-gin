package repository

import (
	"errors"
	"fmt"

	"github.com/go-demo/go-demo-api/application/entity"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	List() []*entity.Customer
	FindById(id string) *entity.Customer
	Update(id string, customer *entity.Customer) error
	Delete(id string) error
}

type customerInMemoryRepository struct {
}

var database []*entity.Customer

func NewCustomerInMemoryRepository() *customerInMemoryRepository {
	database = make([]*entity.Customer, 0)
	return &customerInMemoryRepository{}
}

func (cr *customerInMemoryRepository) List() []*entity.Customer {
	return database
}

func (cr *customerInMemoryRepository) Create(customer *entity.Customer) error {
	database = append(database, customer)
	return nil
}

func (cc *customerInMemoryRepository) FindById(id string) *entity.Customer {
	for _, c := range database {
		if c.ID == id {
			return c
		}
	}

	return nil
}

func (cc *customerInMemoryRepository) Update(id string, customer *entity.Customer) error {
	for i, c := range database {
		if c.ID == id {
			database[i].Name = customer.Name
			return nil
		}
	}

	return fmt.Errorf("Error on update customer: %s", id)
}

func (cc *customerInMemoryRepository) Delete(id string) error {
	for i, c := range database {
		if c.ID == id {
			database = append(database[:i], database[i+1:]...)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Error on delete customer: [%s]", id))
}
