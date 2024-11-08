package main

import (
	"fetch-receipt-processor/internal/handler"
	"fetch-receipt-processor/internal/repo"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//creating points calculator with factory pattern
	pointsCalculator, err := handler.NewPointsCalculatorFactory("one")
	if err != nil {
		fmt.Printf("Error creating points calculator: %v", err)
		return
	}

	repo := repo.NewPointsMap()

	// creating receipt handler using strategy pattern
	recieptHandler := handler.NewReceiptHandler(*repo, pointsCalculator)

	router.POST("/receipts/process", recieptHandler.ProcessReceipt)
	router.GET("/receipts/:id/points", recieptHandler.GetPoints)
	router.Run("0.0.0.0:8080")
}
