package api

import (
	"kalicoin/pkg/api/middlewares"
	"kalicoin/pkg/jaeger"
	"kalicoin/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
	log "github.com/sirupsen/logrus"
)

var tx *pop.Connection

// New creates a new gin instance without starting it
func New(conn *pop.Connection) *gin.Engine {
	tx = conn

	router := gin.Default()

	router.Use(middlewares.TransactionVerification())

	if jaeger.Tracer != nil {
		router.Use(middlewares.Jaeger(jaeger.Tracer))
	} else {
		log.Info("Not using the Jaeger middleware as jaeger isn't initialized")
	}

	router.POST("/payments", payment)
	router.POST("/trades", trade)
	router.POST("/rewards", reward)

	router.GET("/transactions", func(c *gin.Context) {
		var transactions models.Transactions

		if err := tx.All(&transactions); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	})

	router.GET("/payments", func(c *gin.Context) {
		var transactions models.Transactions

		if err := transactions.GetByType(tx, models.Payment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	})

	router.GET("/trades", func(c *gin.Context) {
		var transactions models.Transactions

		if err := transactions.GetByType(tx, models.Trade); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, transactions)
	})

	router.GET("/rewards", func(c *gin.Context) {
		var transactions models.Transactions

		if err := transactions.GetByType(tx, models.Reward); err != nil {
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
