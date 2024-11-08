package main

import (
	"fetch-receipt-processor/internal/handlers"
	"fetch-receipt-processor/internal/repo"
	"fetch-receipt-processor/internal/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Print(err)
		fmt.Printf("Error loading the .env file")
		return
	}

	//creating points calculator with factory pattern
	strategyType := os.Getenv("POINTS_CALC_STRATEGY_TYPE")
	if strategyType == "" {
		strategyType = "one" // default value
	}

	pointsCalculator, err := handlers.NewPointsCalculatorFactory(strategyType)
	if err != nil {
		fmt.Printf("Error creating points calculator: %v", err)
		return
	}

	repo := repo.NewPointsMap()

	// creating receipt handler using strategy pattern
	receiptHandler := handlers.NewReceiptHandler(*repo, pointsCalculator)

	router := routers.InitRouter(receiptHandler)
	router.Run("0.0.0.0:" + os.Getenv("PORT"))
}
