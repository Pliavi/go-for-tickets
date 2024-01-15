package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/pliavi/go-for-tickets/pkg/entities"
)

type ConcertRepository interface {
	Create(name string, totalTickets, concurrentCustomerLimit int) (*entities.Concert, error)
	FindAll(filters ConcertFindAllFilters) ([]*entities.Concert, error)
	FindById(id string) (*entities.Concert, error)
}

type concertRepository struct {
	db *sqlx.DB
}

type ConcertFindAllFilters struct {
	Search string
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

func (cr *concertRepository) FindAll(filters ConcertFindAllFilters) ([]*entities.Concert, error) {
	var rows *sqlx.Rows
	var err error

	query := `SELECT id, name, total_tickets, concurrent_customer_limit FROM concerts`
	if filters.Search == "" {
		rows, err = cr.db.Queryx(query)
	} else {
		query += ` WHERE name ILIKE '%' || $1 || '%' `
		rows, err = cr.db.Queryx(query, filters.Search)
	}

	if err != nil {
		return nil, err
	}

	concerts := []*entities.Concert{}
	for rows.Next() {
		concert := &entities.Concert{}
		err := rows.StructScan(concert)
		if err != nil {
			return nil, err
		}
		concerts = append(concerts, concert)
	}

	return concerts, nil
}

func (cr *concertRepository) FindById(id string) (*entities.Concert, error) {
	query := `SELECT id, name, total_tickets, concurrent_customer_limit FROM concerts WHERE id = $1`
	concert := &entities.Concert{}
	err := cr.db.QueryRowx(query, id).StructScan(concert)
	if err != nil {
		return nil, err
	}

	return concert, nil
}
