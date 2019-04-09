package api

import (
	"net/http"
	"strconv"

	"gitlab.com/bartwillems/kalicoin/pkg/api/middlewares"
	"gitlab.com/bartwillems/kalicoin/pkg/jaeger"
	"gitlab.com/bartwillems/kalicoin/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	log "github.com/sirupsen/logrus"
)

var tx *pop.Connection

// New creates a new gin instance without starting it
func New(conn *pop.Connection) *gin.Engine {
	tx = conn

	router := gin.Default()

	// Basic Auth middleware
	router.Use(gin.BasicAuth(gin.Accounts{
		envy.Get("AUTH_USERNAME", "octaaf"): envy.Get("AUTH_PASSWORD", "secret"),
	}))

	// Transaction middleware, this does error responding of failed transactions
	router.Use(middlewares.TransactionVerification())

	if jaeger.Tracer != nil {
		router.Use(middlewares.Jaeger(jaeger.Tracer))
	} else {
		log.Info("Not using the Jaeger middleware as jaeger isn't initialized")
	}

	router.POST("/payments", payment)
	router.POST("/trades", trade)
	router.POST("/rewards", reward)
	router.POST("/rewards/roll", roll)

	router.GET("/pricetable", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.PriceTable)
	})

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

		err := tx.Transaction(func(tx *pop.Connection) error {
			return tx.All(&wallets)
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, wallets)
	})

	router.GET("/wallets/group/:group_id/owner/:owner_id", func(c *gin.Context) {
		var wallet models.Wallet

		// TODO: use ShouldBindUri to map on wallet directly
		// https://gin-gonic.com/docs/examples/bind-uri/
		groupID, err := strconv.ParseInt(c.Param("group_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Error(err)
			return
		}
		ownerID, err := strconv.Atoi(c.Param("owner_id"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.Error(err)
			return
		}

		err = tx.Transaction(func(tx *pop.Connection) error {
			return wallet.Get(tx, groupID, nulls.NewInt(ownerID))
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, wallet)
	})

	return router
}
