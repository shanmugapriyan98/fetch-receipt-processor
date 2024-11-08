package main

import (
	"fetch-receipt-processor/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/receipts/process", handler.ProcessReceipt)
	router.GET("/receipts/:id/points", handler.GetPoints)
	router.Run("localhost:8080")
}
