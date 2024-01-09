package services

import (
	"log"

	"github.com/pliavi/go-for-tickets/pkg/entities"
	"github.com/pliavi/go-for-tickets/pkg/repositories"
)

type CustomerService interface {
	GetOrCreateCustomer(email string) (*entities.Customer, error)
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}

func (cs *customerService) GetOrCreateCustomer(email string) (*entities.Customer, error) {
	customer, err := cs.customerRepository.GetCustomerByEmail(email)

	if err == nil {
		return customer, nil
	} else {
		log.Print(err)
	}

	customer, err = cs.customerRepository.Create(email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return customer, nil
}
