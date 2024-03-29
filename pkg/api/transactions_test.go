package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/bartwillems/kalicoin/pkg/db"
	"gitlab.com/bartwillems/kalicoin/pkg/models"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/nulls"
	"github.com/stretchr/testify/assert"
)

const transactionAmount = 10
const groupID = 1
const senderID = 1
const receiverID = 2

func Test_Payments(t *testing.T) {
	// Setup the DB
	if err := db.Connect(); err != nil {
		assert.Fail(t, err.Error())
	}

	if err := db.Reset("../../migrations"); err != nil {
		assert.Fail(t, err.Error())
	}

	router := New(db.Conn)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/wallets", nil)
	router.ServeHTTP(w, req)

	// Without basic auth this should fail
	assert.Equal(t, 401, w.Code)

	// Prepare basic auth
	var username = "octaaf"
	var password = "secret"
	envy.Set("AUTH_USERNAME", username)
	envy.Set("AUTH_PASSWORD", password)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/wallets", nil)
	req.SetBasicAuth(username, password)
	router.ServeHTTP(w, req)

	// List wallets, should be empty atm
	var wallets models.Wallets
	walletsJSON, _ := json.Marshal(wallets)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, walletsJSON, w.Body.Bytes())

	// Creating transactions
	w = httptest.NewRecorder()

	trade := models.TradeTransaction{
		GroupID:  groupID,
		Sender:   senderID,
		Receiver: receiverID,
		Amount:   transactionAmount,
	}

	tradeJSON, err := json.Marshal(trade)

	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/trades", bytes.NewBuffer(tradeJSON))

	assert.NoError(t, err)

	req.SetBasicAuth(username, password)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// List wallets after transactions
	w = httptest.NewRecorder()

	wallets = models.Wallets{
		models.Wallet{
			GroupID: groupID,
			OwnerID: senderID,
			Capital: models.StarterCapital - transactionAmount,
		},
		models.Wallet{
			GroupID: groupID,
			OwnerID: receiverID,
			Capital: models.StarterCapital + transactionAmount,
		},
	}
	walletsJSON, _ = json.Marshal(wallets)
	req, _ = http.NewRequest("GET", "/wallets", nil)

	req.SetBasicAuth(username, password)

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var receivedWallets models.Wallets
	err = json.Unmarshal(w.Body.Bytes(), &receivedWallets)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(receivedWallets))

	if len(receivedWallets) != 2 {
		assert.FailNow(
			t,
			fmt.Sprintf("Unable to continue the tests as an invalid wallets response has been given (%v wallets received, 2 expected)", len(receivedWallets)),
		)
	}

	// Order the received wallets for easier testing
	if receivedWallets[0].OwnerID != senderID {
		receivedWallets[0], receivedWallets[1] = receivedWallets[1], receivedWallets[0]
	}

	assert.Equal(t, wallets[0].OwnerID, receivedWallets[0].OwnerID)
	assert.Equal(t, wallets[0].Capital, receivedWallets[0].Capital)

	assert.Equal(t, wallets[1].OwnerID, receivedWallets[1].OwnerID)
	assert.Equal(t, wallets[1].Capital, receivedWallets[1].Capital)

	// Creating transactions
	w = httptest.NewRecorder()

	roll := models.RollReward{
		GroupID:    groupID,
		Receiver:   receiverID,
		Multiplier: nulls.NewUInt32(0),
	}

	rollJSON, err := json.Marshal(roll)

	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/rewards/roll", bytes.NewBuffer(rollJSON))

	assert.NoError(t, err)

	req.SetBasicAuth(username, password)

	router.ServeHTTP(w, req)

	// API call should succeed
	assert.Equal(t, http.StatusCreated, w.Code)

	var transaction models.Transaction
	err = json.Unmarshal(w.Body.Bytes(), &transaction)

	assert.NoError(t, err)
	assert.Equal(t, uint32(20), transaction.Amount)
	assert.Equal(t, "roll - dubs", transaction.Cause.String)
}
