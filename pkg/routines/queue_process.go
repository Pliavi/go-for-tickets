package routines

import (
	"fmt"
	"log"
	"time"

	"github.com/pliavi/go-for-tickets/pkg/models"
)

func QueueProcess() {
	for {
		// each iteration will take 10 seconds
		// TODO: Configurable on environment
		time.Sleep(1 * time.Second)

		concerts, err := models.GetAllConcerts()
		if err != nil {
			log.Fatal(err)
		}

		for _, concert := range concerts {
			queues, err := concert.GetQueues()
			if err != nil {
				log.Fatal(err)
			}

			concertBuyingPhaseLimit := int(concert.BookingSize)
			customersInBuyingPhase := 0
			for _, queue := range queues {
				if customersInBuyingPhase <= concertBuyingPhaseLimit {
					if queue.PurchaseDeadline != nil {
						now := time.Now().UTC()
						if queue.PurchaseDeadline.UTC().Before(now) {
							fmt.Println("=============================================")
							fmt.Println("Purchase deadline.:", *queue.PurchaseDeadline)
							fmt.Println("Current time......:", now)
							fmt.Println("Customer", queue.Customer.ID, "has exceeded the purchase deadline => deleting from queue")

							queue.Delete()
							customersInBuyingPhase--
							continue
						}

						customersInBuyingPhase++
					} else { // customer is not in buying phase
						fmt.Println("=============================================")
						fmt.Println("Customer", queue.Customer.ID, "is in queue => adding to buying phase")
						newDeadline := time.Now().UTC().Add(10 * time.Second)
						fmt.Println("Current time.:", time.Now().UTC())
						fmt.Println("New deadline.:", newDeadline)
						queue.PurchaseDeadline = &newDeadline
						queue.UpdatePurchaseDeadline()
						customersInBuyingPhase++
					}
				} else {
					break
				}
			}

		}
	}
}
