package models

import (
	"kalicoin/pkg/db"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Transaction(t *testing.T) {
	var sender = 1
	var receiver = 2
	var amount uint32 = 10
	var senderWallet, receiverWallet Wallet

	if err := db.Connect(); err != nil {
		assert.Fail(t, err.Error())
	}

    if err := db.Conn.MigrateReset("../../migrations"); err != nil {
		assert.Fail(t, err.Error())
	}

	// The wallets should not yet exist
	err := db.Conn.Where("owner_id = ?", sender).First(&senderWallet)
	assert.Error(t, err)
	err = db.Conn.Where("owner_id = ?", sender).First(&receiverWallet)
	assert.Error(t, err)

	transaction := Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Type:     Trade,
	}

	// This transaction should create both wallets with the correct money
	err = db.Conn.Create(&transaction)
	assert.NoError(t, err)

	// The sender should've lost money
	err = db.Conn.Where("owner_id = ?", sender).First(&senderWallet)
	assert.NoError(t, err)
	assert.Equal(t, StarterCapital-amount, senderWallet.Capital)

	// And the receiver should've gained the money
	err = db.Conn.Where("owner_id = ?", receiver).First(&receiverWallet)
	assert.NoError(t, err)
	assert.Equal(t, StarterCapital+amount, receiverWallet.Capital)

	db.Reset("../../migrations")
}
