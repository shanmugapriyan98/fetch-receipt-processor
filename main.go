package main

import (
	"fetch-receipt-processor/internal/handler"
	"fetch-receipt-processor/internal/repo"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	repo := repo.NewPointsMap()
	recieptHandler := handler.NewReceiptHandler(*repo)
	router.POST("/receipts/process", recieptHandler.ProcessReceipt)
	router.GET("/receipts/:id/points", recieptHandler.GetPoints)
	router.Run("0.0.0.0:8080")
}
