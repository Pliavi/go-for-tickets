package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pliavi/go-for-tickets/pkg/entities"
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
	cookieStr := ""
	cookie, err := r.Cookie("customer_email")
	if err != nil {
		// http.Error(w, "Cookie not found", http.StatusUnauthorized)
		// return

		// generate random email
		cookieStr = ""
	} else {
		cookieStr = cookie.Value
	}

	if cookieStr == "" {
		cookieStr = "customer_" + uuid.NewString() + "@fake.com"
		http.SetCookie(w, cookie)
	}

	customer, err := qh.customerService.GetOrCreateCustomer(cookieStr)
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

func (qh *QueueHandler) PurchaseTicketHandler(w http.ResponseWriter, r *http.Request) {
	// This handler would typically involve more complex logic,
	// such as interacting with a payment gateway and updating ticket availability.

	// For this example, let's assume it marks the customer's purchase as complete.

	// Extract customer information, perform purchase operations...
	customer := r.Context().Value("customer").(entities.Customer)

	// Notify the queueService that the purchase is done
	qh.queueService.PurchaseDone <- &customer

	// Return a success response
	w.WriteHeader(http.StatusOK)
	// Optionally, send back some purchase confirmation details
}
