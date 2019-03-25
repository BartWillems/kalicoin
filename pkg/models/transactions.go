package models

import (
	"errors"
	"reflect"
	"time"

	"github.com/gobuffalo/pop"
	log "github.com/sirupsen/logrus"
)

// TransactionType indicates the payment type
type TransactionType string

// TransactionStatus indicates the current status of a payment
type TransactionStatus string

// Cost is the value of a certain bot feature
type Cost int

// RewardType is the value received by participating in certain bot events
type RewardType uint

// PaymentType is an interface used for purchases
// Every PaymentType should implement Name() to be able to log payment reasons
type PaymentType interface {
	Name() string
}

// Name returns the name of a payment type
func (r RewardType) Name() string {
	val := reflect.Indirect(reflect.ValueOf(r))
	return val.Type().Field(0).Name
}

const (
	// Trade means giving kalicoins from one chat member to another
	Trade TransactionType = "trade"
	// Payment means that a user uses his kalicoins for chat features
	Payment TransactionType = "payment"
	// Reward is money that a user receives for participating in certain bot events
	Reward TransactionType = "reward"

	// Pending indicates that the payment has not been handled
	Pending TransactionStatus = "pending"
	// Succeeded indicates a successfull payment
	Succeeded TransactionStatus = "succeeded"
	// Failed indicates an invalid payment, this could be caused by:
	// - insufficient capital
	// - invalid sender
	// - invalid receiver
	// - general db errors
	Failed TransactionStatus = "failed"

	// Quote (img, video, audio, text)
	Quote Cost = 10

	// All tags everyone with a username and could be annoying
	All Cost = 10

	// Checkin is the biggest money maker as you can't plan participating
	Checkin RewardType = 100

	// Kalivents could be completed twice a day
	Kalivents RewardType = 20
)

// Transaction is a log of a payment that happened
type Transaction struct {
	ID            int64             `json:"id" db:"id"`
	Type          TransactionType   `json:"type" db:"type" binding:"required"`
	Status        TransactionStatus `json:"status" db:"status"`
	Sender        int               `json:"sender" db:"sender" binding:"required"`
	Receiver      int               `json:"receiver" db:"receiver"`
	Amount        uint32            `json:"amount" db:"amount" binding:"required"`
	FailureReason string            `json:"failure_reason" db:"failure_reason"`
	CreatedAt     time.Time         `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at" db:"updated_at"`
}

// Transactions is an array of transactions
type Transactions []Transaction

// BeforeCreate ensures the correct transaction status and wallet config
// This is needed because golang sets the status to ""
// which is not recognized as NULL in postgres
// The transactions have foreign keys to wallets, so a wallet is created
// for the transcation's initiator if the wallet is missing
func (t *Transaction) BeforeCreate(tx *pop.Connection) error {
	t.Status = Pending
	var wallet Wallet

	// Ensure the wallet exists
	return wallet.Get(tx, t.Sender)
}

// AfterCreate will trigger the correct action on a wallet
// This is either a normal payment, or a trade
func (t *Transaction) AfterCreate(tx *pop.Connection) error {
	err := tx.Transaction(func(tx *pop.Connection) error {
		var err error
		switch t.Type {

		case Trade:
			log.Infof("Trading %v kalicoins from %v to %v", t.Amount, t.Sender, t.Receiver)
			err = t.trade(tx)

		case Payment:
			log.Infof("User %v pays %v", t.Sender, t.Amount)
			err = t.pay(tx)

		case Reward:
			log.Infof("Rewarding %v with %v", t.Receiver, t.Amount)
			err = t.receive(tx)

		default:
			err = errors.New("Invalid transaction type")
		}

		if err == nil {
			t.Status = Succeeded
			err = tx.Update(t)
		}

		return err
	})

	if err != nil {
		t.Status = Failed
		t.FailureReason = err.Error()
		return tx.Update(t)
	}

	return nil
}

// Trade takes money from the transaction sender's wallet and adds it to the receiver's wallet
func (t *Transaction) trade(tx *pop.Connection) error {
	if err := t.pay(tx); err != nil {
		return err
	}

	if err := t.receive(tx); err != nil {
		return err
	}

	return nil
}

// Pay takes money from the transaction sender's wallet
func (t *Transaction) pay(tx *pop.Connection) error {
	var wallet Wallet

	if err := wallet.Get(tx, t.Sender); err != nil {
		return err
	}

	if err := wallet.take(t.Amount); err != nil {
		return err
	}

	if err := tx.Update(&wallet); err != nil {
		return err
	}

	return nil
}

// receive is a transaction where a user receives money
func (t *Transaction) receive(tx *pop.Connection) error {
	var wallet Wallet

	if err := wallet.Get(tx, t.Receiver); err != nil {
		return err
	}

	wallet.Capital = wallet.Capital + t.Amount

	return tx.Update(&wallet)
}
