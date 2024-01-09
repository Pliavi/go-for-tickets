package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pliavi/go-for-tickets/pkg/entities"
)

type ConcertRepository interface {
	Create(name string, totalTickets, concurrentCustomerLimit int) (*entities.Concert, error)
	GetConcurrentCustomerLimit() int
}

type concertRepository struct {
	db *sqlx.DB
}

func NewConcertRepository(db *sqlx.DB) ConcertRepository {
	return &concertRepository{
		db: db,
	}
}

func (cr *concertRepository) Create(name string, totalTickets, concurrentCustomerLimit int) (*entities.Concert, error) {
	concert := &entities.Concert{
		Name:                    name,
		TotalTickets:            totalTickets,
		ConcurrentCustomerLimit: concurrentCustomerLimit,
	}

	query := `INSERT INTO concerts (name, total_tickets, concurrent_customer_limit) VALUES ($1, $2, $3) RETURNING id`
	err := cr.db.QueryRowx(query, name, totalTickets, concurrentCustomerLimit).Scan(&concert.ID)
	if err != nil {
		return nil, err
	}

	return concert, nil
}

func (cr *concertRepository) GetConcurrentCustomerLimit() int {
	// TODO: Implement this
	return 5
}
