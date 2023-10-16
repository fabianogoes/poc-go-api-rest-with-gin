package service

import (
	"errors"
	"fmt"

	"github.com/go-demo/go-demo-api/adapter/output/repository"
	"github.com/go-demo/go-demo-api/application/entity"
	"github.com/google/uuid"
)

type CustomerService struct {
	Repository repository.CustomerRepository
}

func NewCustomerService(repository repository.CustomerRepository) *CustomerService {
	return &CustomerService{Repository: repository}
}

func (cs *CustomerService) List() []*entity.Customer {
	return cs.Repository.List()
}

func (cs *CustomerService) Create(name string) error {
	if err := cs.Repository.Create(&entity.Customer{ID: uuid.NewString(), Name: name}); err != nil {
		return errors.New(fmt.Sprintf("Error on create customer: [%s]", name))
	}

	return nil
}

func (cs *CustomerService) FindById(id string) *entity.Customer {
	return cs.Repository.FindById(id)
}

func (cs *CustomerService) Update(id string, customer *entity.Customer) error {
	cs.Repository.Update(id, customer)

	return nil
}

func (cs *CustomerService) Delete(id string) error {
	if err := cs.Repository.Delete(id); err != nil {
		return err
	}

	return nil
}
