package api

import (
	"kalicoin/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pay(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": err.Error()})
		return
	}

	if err := db.Save(&transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"failure_reason": err.Error()})
		return
	}

	if transaction.Status != models.Succeeded {
		c.JSON(http.StatusForbidden, transaction)
	} else {
		c.JSON(http.StatusOK, transaction)
	}
}
