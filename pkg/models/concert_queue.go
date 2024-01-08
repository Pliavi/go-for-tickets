package models

import (
	"time"

	"github.com/pliavi/go-for-tickets/pkg/utils/database"
)

type ConcertQueue struct {
	ID               *int // TODO: use uuid
	Concert          *Concert
	Customer         *Customer
	PurchaseDeadline *time.Time
}

func NewConcertQueue(concert *Concert, customer *Customer) *ConcertQueue {
	return &ConcertQueue{
		Concert:  concert,
		Customer: customer,
	}
}

func (cq *ConcertQueue) Save() error {
	db := database.GetInstance()

	var lid int
	err := db.
		QueryRow("INSERT INTO concert_queues (concert_id, customer_id) VALUES ($1, $2) RETURNING id", cq.Concert.ID, cq.Customer.ID).
		Scan(&lid)

	if err != nil {
		panic(err)
		// return err
	}

	if err != nil {
		panic(err)
		// return err
	}

	cq.ID = &lid

	return nil
}

func (cq *ConcertQueue) UpdatePurchaseDeadline() error {
	db := database.GetInstance()

	_, err := db.
		Exec("UPDATE concert_queues SET purchase_deadline = $1 WHERE id = $2", cq.PurchaseDeadline, cq.ID)

	if err != nil {
		return err
	}

	return nil
}

// TODO: This is not complete
//
// The estimation calculation is based on:
// - the number of people in front of the customer
//   - if there are slots available, the customer will be
//     placed at the buying process/phases instantly(estimation <= 0)
//   - this number will be divided by the booking size
//   - the result will be multiplied by the mean duration of
//     the concert time buying process/phases
//
// example:
// - concert booking size: 10
// - customers in queue: 20
// - mean duration of concert time buying process: 5 minutes(TODO: needs calculation)
// - estimated time in queue: (20-10) / 10 * 5 = 10 minutes
// NOTE: This estimation will not be saved in the database(maybe it should?)
func (cq *ConcertQueue) EstimatedTimeInQueue() (*int, error) {
	db := database.GetInstance()

	sql := `
		SELECT COUNT(*)
		FROM concert_queues
		WHERE concert_id = $1 AND id < $2
	`
	var queueSize int
	err := db.
		QueryRow(
			sql,
			cq.Concert.ID,
			cq.ID,
		).
		Scan(&queueSize)

	// TODO: dividing ints in Go will round or truncate the result?
	estimatedTimeInSecs := queueSize / int(cq.Concert.BookingSize) * 5

	if err != nil {
		return nil, err
	}

	return &estimatedTimeInSecs, nil
}

func GetConcertQueue(concert_id string, customer_id string) (*ConcertQueue, error) {
	db := database.GetInstance()

	var concert Concert
	err := db.
		QueryRow("SELECT id, name, booking_size FROM concerts WHERE id::text = $1", concert_id).
		Scan(&concert.ID, &concert.Name, &concert.BookingSize)

	if err != nil {
		return nil, err
	}

	var customer Customer
	err = db.
		QueryRow("SELECT id FROM customers WHERE id::text = $1", customer_id).
		Scan(&customer.ID)

	if err != nil {
		return nil, err
	}

	var concert_queue ConcertQueue
	err = db.
		QueryRow("SELECT id FROM concert_queues WHERE concert_id = $1 AND customer_id::text = $2",
			concert.ID,
			customer.ID,
		).
		Scan(&concert_queue.ID)

	if err != nil {
		return nil, err
	}

	concert_queue.Concert = &concert
	concert_queue.Customer = &customer

	return &concert_queue, nil
}

func (cq *ConcertQueue) Delete() error {
	db := database.GetInstance()

	_, err := db.
		Exec("DELETE FROM concert_queues WHERE id = $1", cq.ID)

	if err != nil {
		return err
	}

	return nil
}
