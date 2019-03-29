package middlewares

import (
	"errors"
	"kalicoin/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// TransactionVerification checks if the request returns a transaction
// If this is the case, it will check if the transaction was succesfull
// and returns the correct response to the client
func TransactionVerification() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		transaction, exists := isTransaction(c.Get("transaction"))

		if exists {
			log.Info("Transaction exists")
			if transaction.Status != models.Succeeded {
				c.Error(errors.New(transaction.FailureReason.String))
				c.JSON(http.StatusForbidden, transaction)
			} else {
				c.JSON(http.StatusCreated, transaction)
			}
		} else {
			log.Info("Transaction does not exist")
		}
	}
}

func isTransaction(t interface{}, exists bool) (*models.Transaction, bool) {
	if !exists {
		return &models.Transaction{}, false
	}

	switch v := t.(type) {
	case *models.Transaction:
		return v, true
	default:
		return &models.Transaction{}, false
	}
}
