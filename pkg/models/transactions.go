package models

import (
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
)

type Transaction struct {
	ID        int64     `json:"id" db:"id"`
	Sender    int       `json:"sender" db:"sender"`
	Receiver  int       `json:"receiver" db:"receiver"`
	Amount    int       `json:"amount" db:"amount"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// Transactions is an array of transactions
type Transactions []Transaction

// ValidateCreate should check if the paymend could complete.
// a payment may fail if the amount is < 0 or if one of sender/receiver does not exist
func (t *Transaction) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
