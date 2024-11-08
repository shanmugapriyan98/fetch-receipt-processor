package handler

import (
	"fetch-receipt-processor/pkg/models"
	"fetch-receipt-processor/pkg/utils"
)

func calculatePoints(receipt models.Receipt) int64 {
	points := int64(0)

	points += utils.CalculateRetailerRewards(receipt.Retailer)
	// fmt.Println("Points afer 1:", points)
	points += utils.CalculateWholeAmountRewards(receipt.Total)
	// fmt.Println("Points afer 2:", points)
	points += utils.CalculateMultipleOf25Rewards(receipt.Total)
	// fmt.Println("Points afer 3:", points)
	points += utils.CalculateDoubleItemRewards(receipt.Items)
	// fmt.Println("Points afer 4:", points)
	points += utils.CalculateItemDescRewards(receipt.Items)
	// fmt.Println("Points afer 5:", points)
	points += utils.CalculatePurchaseDateRewards(receipt.PurchaseDate)
	// fmt.Println("Points afer 6:", points)
	points += utils.CalculatePurchaseTimeRewards(receipt.PurchaseTime)
	// fmt.Println("Points afer 7:", points)

	return points
}
