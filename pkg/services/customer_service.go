package services

import (
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
	}

	customer, err = cs.customerRepository.Create(email)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
