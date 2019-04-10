package api

import (
	"bytes"
	"encoding/json"
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

	// Create first wallet
	w = httptest.NewRecorder()

	wallet := models.Wallet{
		GroupID: groupID,
		OwnerID: senderID,
	}

	walletJSON, err := json.Marshal(wallet)

	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/wallets", bytes.NewBuffer(walletJSON))

	assert.NoError(t, err)

	req.SetBasicAuth(username, password)

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	// Creating transactions
	w = httptest.NewRecorder()

	payment := models.PaymentTransaction{
		GroupID: groupID,
		Sender:  senderID,
		Cause:   nulls.NewString("roll"),
	}

	paymentJSON, err := json.Marshal(payment)

	assert.NoError(t, err)

	req, err = http.NewRequest("POST", "/payments", bytes.NewBuffer(paymentJSON))

	assert.NoError(t, err)

	req.SetBasicAuth(username, password)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	// Creating transactions
	w = httptest.NewRecorder()

	roll := models.RollReward{
		GroupID:    groupID,
		Receiver:   senderID,
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
