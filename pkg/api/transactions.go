package api

import (
	"errors"
	"net/http"

	"gitlab.com/bartwillems/kalicoin/pkg/models"

	"github.com/gin-gonic/gin"
)

func payment(c *gin.Context) {
	var payment models.PaymentTransaction

	if err := c.ShouldBindJSON(&payment); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": err.Error()})
		return
	}

	transaction, err := payment.Create(tx)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"failure_reason": err.Error()})
		c.Error(err)
		return
	}

	c.Set("transaction", transaction)
}

func trade(c *gin.Context) {
	var trade models.TradeTransaction

	if err := c.ShouldBindJSON(&trade); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": err.Error()})
		return
	}

	if trade.Sender == trade.Receiver {
		c.Error(errors.New("Sender is receiver"))
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": "You can not send money to yourself"})
		return
	}

	transaction, err := trade.Create(tx)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusForbidden, gin.H{"failure_reason": err.Error()})
		return
	}

	c.Set("transaction", transaction)
}

func reward(c *gin.Context) {
	var reward models.RewardTransaction

	if err := c.ShouldBindJSON(&reward); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": err.Error()})
		return
	}

	transaction, err := reward.Create(tx)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"failure_reason": err.Error()})
		return
	}

	c.Set("transaction", transaction)
}
