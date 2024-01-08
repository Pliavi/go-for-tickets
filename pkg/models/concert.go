package models

import (
	"errors"
	"strings"

	"github.com/pliavi/go-for-tickets/pkg/utils/database"
)

type Concert struct {
	ID          *int
	Name        string
	BookingSize uint
}

func NewConcert(name string, bookingSize uint) (*Concert, error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("name cannot be empty")
	}
	if bookingSize == 0 {
		return nil, errors.New("booking size cannot be zero")
	}

	c := &Concert{
		Name:        name,
		BookingSize: bookingSize,
	}

	return c, nil
}

func (c *Concert) Save() error {
	db := database.GetInstance()
	res, err := db.Exec(
		"INSERT INTO concerts (name, booking_size) VALUES ($1, $2)  RETURNING id",
		c.Name,
		c.BookingSize,
	)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// TODO: maybe there is a better way to do this?
	//       why i need a new variable to get the reference?
	//			 why i can't just do `c.id = &int(lid)`?
	ilid := int(lid)
	c.ID = &ilid

	return nil
}

func GetAllConcerts() ([]*Concert, error) {
	db := database.GetInstance()

	rows, err := db.Query("SELECT id, name, booking_size FROM concerts")
	if err != nil {
		return nil, err
	}

	var concerts []*Concert
	for rows.Next() {
		var concert Concert
		err := rows.Scan(
			&concert.ID,
			&concert.Name,
			&concert.BookingSize,
		)
		if err != nil {
			return nil, err
		}

		concerts = append(concerts, &concert)
	}

	return concerts, nil
}

func GetConcert(id string) (*Concert, error) {
	db := database.GetInstance()

	var concert Concert
	err := db.
		QueryRow(
			"SELECT id, name, booking_size FROM concerts WHERE id = $1",
			id,
		).
		Scan(
			&concert.ID,
			&concert.Name,
			&concert.BookingSize,
		)

	if err != nil {
		return nil, err
	}

	return &concert, nil
}

func (c *Concert) GetQueues() ([]*ConcertQueue, error) {
	db := database.GetInstance()

	rows, err := db.
		Query(
			"SELECT id, customer_id, purchase_deadline FROM concert_queues WHERE concert_id = $1 ORDER BY id ASC LIMIT $2",
			c.ID,
			c.BookingSize*2,
		)

	if err != nil {
		return nil, err
	}

	var queues []*ConcertQueue

	for rows.Next() {
		var queue ConcertQueue
		queue.Concert = c
		queue.Customer = &Customer{}
		err := rows.Scan(
			&queue.ID,
			&queue.Customer.ID,
			&queue.PurchaseDeadline,
		)
		if err != nil {
			return nil, err
		}

		queue.Concert = c
		queues = append(queues, &queue)
	}

	return queues, nil
}
