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
	startPurchase chan entities.Customer
	endPurchase   chan entities.Customer
}

func NewQueueService(concertRepo repositories.ConcertRepository, customerRepo repositories.CustomerRepository) *QueueService {
	return &QueueService{
		concertRepo:   concertRepo,
		customerRepo:  customerRepo,
		queue:         []entities.Customer{},
		startPurchase: make(chan entities.Customer),
		endPurchase:   make(chan entities.Customer),
	}
}
func (qs *QueueService) AddCustomerToQueue(customer *entities.Customer) (time.Duration, error) {
	qs.queueLock.Lock()
	defer qs.queueLock.Unlock()

	estimatedWaitTime := time.Duration(0)

	if qs.canMoveToBuyingPhase() {
		// If a spot is available, send the customer to the startPurchase channel
		qs.startPurchase <- *customer
	} else {
		// If no spot is available, add the customer to the queue
		qs.queue = append(qs.queue, *customer)
		// Calculate the estimated wait time
		estimatedWaitTime = qs.calculateEstimatedWaitTime()
	}

	return estimatedWaitTime, nil
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
