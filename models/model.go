package models

import (
	"time"
)

type UserTransaction struct {
	ID            int `json:",omitempty"`
	UserID        string
	Amount        float64
	TransactionID string    `json:",omitempty"`
	Timestamp     time.Time `json:",omitempty"`
}
