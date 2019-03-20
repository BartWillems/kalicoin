package api

import (
	"kalicoin/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/pop"
)

var db *pop.Connection

// New creates a new gin instance without starting it
func New(conn *pop.Connection) *gin.Engine {
	db = conn

	router := gin.Default()

	router.POST("/payment", pay)

	return router
}

func pay(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Transaction successfully stored"})
}
