package models

import (
	"testing"

	"gitlab.com/bartwillems/kalicoin/pkg/db"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/stretchr/testify/assert"
)

func Test_Transaction(t *testing.T) {
	var groupID int64 = 1
	var sender = nulls.NewInt(1)
	var receiver = nulls.NewInt(2)
	var senderWallet, receiverWallet Wallet

	if err := db.Connect(); err != nil {
		assert.Fail(t, err.Error())
	}

	if err := db.Reset("../../migrations"); err != nil {
		assert.Fail(t, err.Error())
	}

	senderWallet = Wallet{
		GroupID: groupID,
		OwnerID: sender.Int,
	}
	err := senderWallet.Create(db.Conn)
	assert.NoError(t, err)

	receiverWallet = Wallet{
		GroupID: groupID,
		OwnerID: receiver.Int,
	}
	err = receiverWallet.Create(db.Conn)
	assert.NoError(t, err)

	// TODO: uncomment trades: currently disabled as trading is not yet implemented
	// trade := TradeTransaction{
	// 	GroupID:  groupID,
	// 	Sender:   sender.Int,
	// 	Receiver: receiver.Int,
	// 	Amount:   amount,
	// }

	// transaction, err := trade.Create(db.Conn)

	// // This transaction should create both wallets with the correct money
	// // err = db.Conn.Create(&transaction)
	// assert.NoError(t, err)
	// assert.Equal(t, Succeeded, transaction.Status)
	// assert.Equal(t, Trade, transaction.Type)

	// // The sender should've lost money
	// err = db.Conn.Transaction(func(tx *pop.Connection) error {
	// 	return senderWallet.Get(tx, groupID, sender)
	// })
	// assert.NoError(t, err)
	// assert.Equal(t, StarterCapital-amount, senderWallet.Capital)

	// // And the receiver should've gained the money
	// err = db.Conn.Transaction(func(tx *pop.Connection) error {
	// 	return receiverWallet.Get(tx, groupID, receiver)
	// })
	// assert.NoError(t, err)
	// assert.Equal(t, StarterCapital+amount, receiverWallet.Capital)

	// Test rewards
	reward := RewardTransaction{
		GroupID:  groupID,
		Receiver: receiver.Int,
		Cause:    nulls.NewString("checkin"),
	}

	// Store the receiver's old capital
	capital := receiverWallet.Capital

	rewardTransaction, err := reward.Create(db.Conn)

	assert.NoError(t, err)
	assert.Equal(t, Succeeded, rewardTransaction.Status)
	assert.Equal(t, Reward, rewardTransaction.Type)

	err = db.Conn.Transaction(func(tx *pop.Connection) error {
		return receiverWallet.Get(tx, groupID, receiver)
	})
	assert.NoError(t, err)
	assert.Equal(t, capital+PriceTable[Reward]["checkin"], receiverWallet.Capital)
}
