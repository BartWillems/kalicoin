package api

import (
	"kalicoin/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
)

var tx *pop.Connection

// New creates a new gin instance without starting it
func New(conn *pop.Connection) *gin.Engine {
	tx = conn

	router := gin.Default()

	router.POST("/transactions", pay)

	router.GET("/transactions", func(c *gin.Context) {
		var transactions models.Transactions

		if err := tx.All(&transactions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	})

	router.GET("/wallets", func(c *gin.Context) {
		var wallets models.Wallets

		if err := tx.All(&wallets); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, wallets)
	})

	return router
}
