package services

import (
	"fmt"
	"sync"
	"time"

	"github.com/pliavi/go-for-tickets/pkg/entities"
	"github.com/pliavi/go-for-tickets/pkg/repositories"
)

type QueueService struct {
	concertRepo   repositories.ConcertRepository
	customerRepo  repositories.CustomerRepository
	queue         []entities.Customer
	queueLock     sync.Mutex
	StartPurchase chan *entities.Customer
	PurchaseDone  chan *entities.Customer
}

func NewQueueService(concertRepo repositories.ConcertRepository, customerRepo repositories.CustomerRepository) *QueueService {
	return &QueueService{
		concertRepo:   concertRepo,
		customerRepo:  customerRepo,
		queue:         []entities.Customer{},
		StartPurchase: make(chan *entities.Customer),
		PurchaseDone:  make(chan *entities.Customer),
	}
}

func (qs *QueueService) AddCustomerToQueue(customer *entities.Customer) (time.Duration, error) {
	qs.queueLock.Lock()
	defer qs.queueLock.Unlock()

	estimatedWaitTime := time.Duration(0)

	if qs.canMoveToBuyingPhase() {
		// If a spot is available, send the customer to the startPurchase channel
		qs.StartPurchase <- customer
	} else {
		// If no spot is available, add the customer to the queue
		qs.queue = append(qs.queue, *customer)
		// Calculate the estimated wait time
		estimatedWaitTime = qs.calculateEstimatedWaitTime()
	}

	return estimatedWaitTime, nil
}

func (qs *QueueService) ProcessQueue() {
	for {
		select {
		case customer := <-qs.StartPurchase:
			// Customer is ready to start purchasing
			// Handle logic to initiate purchase
			go qs.initiatePurchase(customer)

		case customer := <-qs.PurchaseDone:
			// Customer has completed or failed the purchase
			// Handle post-purchase logic
			qs.handlePostPurchase(customer)
		}

		// Add logic to manage adding/removing from the queue
	}
}

func (qs *QueueService) initiatePurchase(customer *entities.Customer) {
	// Set a timer for the purchase duration
	purchaseTimeLimit := time.Duration(qs.concertRepo.GetConcurrentCustomerLimit())
	timer := time.NewTimer(purchaseTimeLimit)

	<-timer.C // is this necessary?

	// Time expired, send customer to purchaseDone channel
	qs.PurchaseDone <- customer
	// Add logic to handle customer completing the purchase before the timeout
}

func (qs *QueueService) handlePostPurchase(customer *entities.Customer) {
	// Remove the customer from the queue and/or active buyers
	// Update the customer's status or perform other cleanup actions
	qs.queueLock.Lock()
	defer qs.queueLock.Unlock()

	// Remove the customer from the queue
	for i, c := range qs.queue {
		if c.ID == customer.ID {
			qs.queue = append(qs.queue[:i], qs.queue[i+1:]...)
			break
		}
	}
}

func (qs *QueueService) canMoveToBuyingPhase() bool {
	// Implement the logic to check if a customer can move to the buying phase
	// For example, check if the number of customers currently buying is less than the limit
	return len(qs.queue) < qs.concertRepo.GetConcurrentCustomerLimit()
}

func (qs *QueueService) calculateEstimatedWaitTime() time.Duration {
	// Implement the logic to calculate the estimated wait time
	// This might depend on the queue length and average processing time per customer
	averageProcessingTimePerCustomer := 5 * time.Minute

	return time.Duration(len(qs.queue)) * averageProcessingTimePerCustomer
}

func (qs *QueueService) notifyCustomerWaitTime(customer *entities.Customer, waitTime time.Duration) {
	// Implement the logic to notify the customer about their estimated wait time
	// This could be sending an email, a message, or updating a status somewhere

	// For now, just print the wait time to the console
	fmt.Println("Estimated wait time for customer", customer.ID, "is", waitTime)
}
