package models

import (
	"github.com/google/uuid"

	"github.com/pliavi/go-for-tickets/pkg/utils/database"
)

type Customer struct {
	ID uuid.UUID
}

func NewCustomer() *Customer {
	return &Customer{
		ID: uuid.New(),
	}
}

func (c *Customer) Save() error {
	db := database.GetInstance()

	var lid uuid.UUID
	err := db.
		// TODO: create using default
		QueryRow("INSERT INTO customers (id) VALUES ($1) RETURNING id", c.ID).
		Scan(&lid)

	if err != nil {
		return err
	}

	return nil
}

func GetCustomer(id string) (*Customer, error) {
	db := database.GetInstance()

	var customer Customer
	err := db.QueryRow("SELECT id FROM customers WHERE id::text = $1", id).Scan(&customer.ID)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
