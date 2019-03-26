package models

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
)

// RewardTransaction is a struct used for creating rewards
type RewardTransaction struct {
	*baseTransaction
	Receiver int          `json:"receiver" db:"receiver" binding:"required"`
	Cause    nulls.String `json:"cause" db:"cause" binding:"required"`
}

// Create maps a reward to a real transaction and creates it
func (r *RewardTransaction) Create(tx *pop.Connection) (*Transaction, error) {
	var wallet Wallet

	if err := wallet.Get(tx, nulls.NewInt(r.Receiver)); err != nil {
		return nil, err
	}

	transaction := Transaction{
		Receiver: nulls.NewInt(r.Receiver),
		Cause:    r.Cause,
		Type:     Reward,
	}
	err := tx.Create(&transaction)
	return &transaction, err
}
