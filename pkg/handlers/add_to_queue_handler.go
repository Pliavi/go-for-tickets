package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pliavi/go-for-tickets/pkg/services"
)

type QueueHandler struct {
	queueService    services.QueueService
	customerService services.CustomerService
}

func NewQueueHandler(queueService services.QueueService, customerService services.CustomerService) *QueueHandler {
	return &QueueHandler{
		queueService:    queueService,
		customerService: customerService,
	}
}

func (qh *QueueHandler) AddCustomerToQueueHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("customer_email")
	if err != nil {
		http.Error(w, "Cookie not found", http.StatusUnauthorized)
		return
	}

	customer, err := qh.customerService.GetOrCreateCustomer(cookie.Value)
	if err != nil {
		http.Error(w, "Failed to retrieve or create customer", http.StatusInternalServerError)
		return
	}

	waitTime, err := qh.queueService.AddCustomerToQueue(customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]time.Duration{"waitTime": waitTime})
}
