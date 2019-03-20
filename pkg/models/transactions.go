package models

import (
	"time"
)

// TransactionType indicates the payment type
type TransactionType string

const (
	// Trade means giving kalicoins from one chat member to another
	Trade TransactionType = "trade"
	// Payment means that a user uses his kalicoins for chat features
	Payment TransactionType = "payment"
)

// Transaction is a log of a payment that happened
type Transaction struct {
	ID        int64           `json:"id" db:"id"`
	Type      TransactionType `json:"type" db:"type" binding:"required"`
	Sender    int             `json:"sender" db:"sender" binding:"required"`
	Receiver  int             `json:"receiver" db:"receiver"`
	Amount    uint32          `json:"amount" db:"amount" binding:"required"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

// Transactions is an array of transactions
type Transactions []Transaction
