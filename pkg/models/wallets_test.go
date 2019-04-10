package models

import (
	"testing"

	"gitlab.com/bartwillems/kalicoin/pkg/db"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/stretchr/testify/assert"
)

func Test_Wallet(t *testing.T) {
	var groupID int64 = 1
	var userID = 10
	var smallPayment = StarterCapital / 10
	var wallet Wallet

	if err := db.Connect(); err != nil {
		assert.Fail(t, err.Error())
	}

	if err := db.Reset("../../migrations"); err != nil {
		assert.Fail(t, err.Error())
	}

	pop.Debug = true

	// Fetch a new user's wallet
	err := db.Conn.Where("owner_id = ?", userID).First(&wallet)

	// Wallet should not yet exist
	assert.Error(t, err)

	// wallet.Get should no longer create a wallet
	err = wallet.Get(db.Conn, groupID, nulls.NewInt(userID))

	assert.Error(t, err)

	wallet.OwnerID = userID
	wallet.GroupID = groupID

	err = wallet.Create(db.Conn)

	// The wallet should be created
	assert.NoError(t, err)

	// A new wallet should start with the starter capital
	assert.Equal(t, StarterCapital, wallet.Capital)

	// Taking a small amount of money from the wallet should work
	err = wallet.pay(db.Conn, smallPayment)
	assert.NoError(t, err)
	wallet.Get(db.Conn, groupID, nulls.NewInt(userID))
	assert.EqualValues(t, StarterCapital-smallPayment, wallet.Capital)

	// It should not be possible to take more money than what is left
	err = wallet.pay(db.Conn, wallet.Capital+1)
	assert.Error(t, err)

	// But it should be possible to empty out your wallet!
	err = wallet.pay(db.Conn, StarterCapital-smallPayment)
	assert.NoError(t, err)
}
