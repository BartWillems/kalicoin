package api

import (
	"net/http"

	"gitlab.com/bartwillems/kalicoin/pkg/models"

	"github.com/gin-gonic/gin"
)

func wallet(c *gin.Context) {
	var wallet models.Wallet

	if err := c.ShouldBindJSON(&wallet); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"failure_reason": err.Error()})
		return
	}

	err := wallet.Create(tx)

	if err != nil {
		c.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"failure_reason": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, wallet)
}
