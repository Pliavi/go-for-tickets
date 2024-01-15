package entities

type Concert struct {
	ID                      *uint  `json:"id"`
	Name                    string `json:"name"`
	TotalTickets            int    `json:"total_tickets" db:"total_tickets"`
	ConcurrentCustomerLimit int    `json:"concurrent_customer_limit" db:"concurrent_customer_limit"`
}

func NewConcert(name string, totalTickets, concurrentCustomerLimit int) *Concert {
	return &Concert{
		Name:                    name,
		TotalTickets:            totalTickets,
		ConcurrentCustomerLimit: concurrentCustomerLimit,
	}
}
