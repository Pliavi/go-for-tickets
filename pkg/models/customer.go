package models

// import (
// 	"time"

// 	"github.com/google/uuid"
// )

// type Customer struct {
// 	Id                   uuid.UUID  `json:"id"`
// 	SessionEnd           *time.Time `json:"session_end"`
// 	EstimatedTimeInQueue *int       `json:"estimated_time_in_queue"`
// }

// func NewCustomer(
// 	id *uuid.UUID,
// 	sessionEnd time.Time,
// 	estimatedTimeInQueue int,
// ) *Customer {
// 	if id == nil {
// 		newId := uuid.New()
// 		id = &newId
// 	}

// 	return &Customer{
// 		Id:                   *id,
// 		SessionEnd:           &sessionEnd,
// 		EstimatedTimeInQueue: &estimatedTimeInQueue,
// 	}
// }

// func (c *Customer) SetSessionEnd(t time.Time) *Customer {
// 	c.SessionEnd = &t
// 	return c
// }

// func (c *Customer) SetEstimatedTimeInQueue(t int) *Customer {
// 	if c.SessionEnd != nil {
// 		c.EstimatedTimeInQueue = &t
// 		return c
// 	}

// 	panic("Cannot set estimated time after customer session started, are you trying to return the customer to the queue?")
// }

// func (c *Customer) IsSessionFinished() bool {
// 	return time.Now().Compare(*c.SessionEnd) >= 0
// }
