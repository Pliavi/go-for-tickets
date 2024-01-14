package services

import (
	"github.com/pliavi/go-for-tickets/pkg/repositories"
)

type CustomerService interface {
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
