package main

import (
	"fetch-receipt-processor/internal/handlers"
	"fetch-receipt-processor/internal/repo"
	"fetch-receipt-processor/internal/routers"
	"fmt"
)

func main() {

	//creating points calculator with factory pattern
	pointsCalculator, err := handlers.NewPointsCalculatorFactory("one")
	if err != nil {
		fmt.Printf("Error creating points calculator: %v", err)
		return
	}
	repo := repo.NewPointsMap()

	// creating receipt handler using strategy pattern
	receiptHandler := handlers.NewReceiptHandler(*repo, pointsCalculator)

	router := routers.InitRouter(receiptHandler)
	router.Run("0.0.0.0:8080")
}
