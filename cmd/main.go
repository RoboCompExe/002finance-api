package main

import (
	"finance-api/internal/database"
	"finance-api/internal/handler"
	"finance-api/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	svc := &service.TransferService{DB: database.DB}
	h := &handler.TransferHandler{Service: svc}

	r.POST("/transfer", h.Transfer)

	r.Run(":8080")
}
