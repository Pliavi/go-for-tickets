package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pliavi/go-for-tickets/pkg/entities"
)

type CustomerRepository interface {
	Create(name string) (*entities.Customer, error)
	GetCustomerByEmail(email string) (*entities.Customer, error)
}

type customerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (cr *customerRepository) Create(email string) (*entities.Customer, error) {
	customer := &entities.Customer{}

	query := `INSERT INTO customers (name) VALUES ($1) RETURNING id`
	err := cr.db.QueryRowx(query, email).StructScan(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (cr *customerRepository) GetCustomerByEmail(email string) (*entities.Customer, error) {
	customer := &entities.Customer{}

	query := `SELECT id, name FROM customers WHERE email = $1`
	err := cr.db.QueryRowx(query, email).StructScan(customer)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
