package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/pliavi/go-for-tickets/pkg/models"
	"github.com/pliavi/go-for-tickets/pkg/utils"
	"github.com/pliavi/go-for-tickets/pkg/utils/database"
)

func AddToQueueHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetInstance()
	tx, err := db.Begin()

	if err != nil {
		utils.SendDefaultErrorResponse(w, err)
		return
	}

	concert_id := r.URL.Query().Get("id")
	if concert_id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("concert id is required"))
		return
	}

	var customer_id string
	if len(r.Cookies()) > 0 {
		customer_id = r.Cookies()[0].Value
	}

	var customer *models.Customer
	var concert_queue *models.ConcertQueue

	if customer_id == "" {
		customer, concert_queue, err = createCustomerAndQueue(concert_id)
	} else {
		customer, concert_queue, err = getCustomerAndQueue(concert_id, customer_id)
	}
	if err != nil {
		utils.SendDefaultErrorResponse(w, err)
		return
	}

	estimated_queue_duration, err := concert_queue.EstimatedTimeInQueue()
	if err != nil {
		utils.SendDefaultErrorResponse(w, err)
		return
	}

	err = tx.Commit()
	if err != nil {
		utils.SendDefaultErrorResponse(w, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "customer_id",
		Value: customer.ID.String(),
	})

	utils.SendJsonResponse(
		w,
		http.StatusCreated,
		map[string]string{ // Why did copilot suggest map[string]interface?
			"queue_id":                 strconv.Itoa(*concert_queue.ID),
			"estimated_queue_duration": fmt.Sprintf("%d seconds", *estimated_queue_duration),
		},
	)
}

func createCustomerAndQueue(concert_id string) (*models.Customer, *models.ConcertQueue, error) {
	customer := models.NewCustomer()

	err := customer.Save()
	if err != nil {
		return nil, nil, err
	}

	concert, err := models.GetConcert(concert_id)
	if err != nil {
		return nil, nil, err
	}

	concert_queue := models.NewConcertQueue(concert, customer)
	err = concert_queue.Save()
	if err != nil {
		return nil, nil, err
	}

	return customer, concert_queue, nil
}

func getCustomerAndQueue(concert_id string, customer_id string) (*models.Customer, *models.ConcertQueue, error) {
	customer, err := models.GetCustomer(customer_id)
	if err != nil {
		return nil, nil, err
	}

	concert_queue, err := models.GetConcertQueue(concert_id, customer_id)
	if err != nil && err.Error() == "sql: no rows in result set" {
		concert, err := models.GetConcert(concert_id)
		if err != nil {
			return nil, nil, err
		}

		concert_queue = models.NewConcertQueue(concert, customer)
		err = concert_queue.Save()
		if err != nil {
			return nil, nil, err
		}
	}

	return customer, concert_queue, nil
}
