package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"finance-api/internal/service"
)

type TransferHandler struct {
	Service *service.TransferService
}

func (h *TransferHandler) Transfer(c *gin.Context) {
	var req struct {
		From   string `json:"from"`
		To     string `json:"to"`
		Amount int64  `json:"amount"`
	}

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.Transfer(c, req.From, req.To, req.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
