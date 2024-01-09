package entities

type Concert struct {
	ID                      *uint
	Name                    string
	TotalTickets            int
	ConcurrentCustomerLimit int
}

func NewConcert(name string, totalTickets, concurrentCustomerLimit int) *Concert {
	return &Concert{
		Name:                    name,
		TotalTickets:            totalTickets,
		ConcurrentCustomerLimit: concurrentCustomerLimit,
	}
}
