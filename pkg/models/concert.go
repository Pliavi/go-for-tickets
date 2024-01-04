package models

import (
	"errors"
	"strings"

	"github.com/pliavi/go-for-tickets/pkg/utils/database"
)

type Concert struct {
	id          *int
	name        string
	bookingSize uint
}

func NewConcert(name string, bookingSize uint) (*Concert, error) {
	if strings.Trim(name, " ") == "" {
		return nil, errors.New("name cannot be empty")
	}
	if bookingSize == 0 {
		return nil, errors.New("booking size cannot be zero")
	}

	c := &Concert{
		name:        name,
		bookingSize: bookingSize,
	}

	return c, nil
}

func (c *Concert) Save() error {
	db := database.GetInstance()
	res, err := db.Exec("INSERT INTO concerts (name, booking_size) VALUES ($1, $2)", c.name, c.bookingSize)
	if err != nil {
		return err
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// TODO: maybe there is a better way to do this?
	//       why i need a new variable to get the reference?
	//			 why i can't just do `c.id = &int(lid)`?
	ilid := int(lid)
	c.id = &ilid

	return nil
}
