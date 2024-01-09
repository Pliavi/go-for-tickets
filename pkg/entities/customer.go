package entities

type Customer struct {
	ID    *uint  `json:"id"`
	Email string `json:"email"`
}

func NewCustomer(ID *uint, email string) *Customer {
	return &Customer{
		ID:    ID,
		Email: email,
	}
}
