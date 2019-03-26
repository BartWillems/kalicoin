package models

import (
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
)

// TradeTransaction is a struct used for creating trades
type TradeTransaction struct {
	*baseTransaction
	Sender   int    `json:"sender" db:"sender" binding:"required"`
	Receiver int    `json:"receiver" db:"receiver" binding:"required"`
	Amount   uint32 `json:"amount" db:"amount" binding:"required"`
}

// Create maps a trade to a real transaction and creates it
func (t *TradeTransaction) Create(tx *pop.Connection) (*Transaction, error) {
	var senderWallet, receiverWallet Wallet

	if err := senderWallet.Get(tx, nulls.NewInt(t.Sender)); err != nil {
		return nil, err
	}

	if err := receiverWallet.Get(tx, nulls.NewInt(t.Receiver)); err != nil {
		return nil, err
	}

	transaction := Transaction{
		Sender:   nulls.NewInt(t.Sender),
		Receiver: nulls.NewInt(t.Receiver),
		Amount:   t.Amount,
		Type:     Trade,
	}
	err := tx.Create(&transaction)
	return &transaction, err
}