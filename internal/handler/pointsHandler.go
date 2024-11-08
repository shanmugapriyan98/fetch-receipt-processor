package handler

import (
	"fetch-receipt-processor/pkg/models"
	"fetch-receipt-processor/pkg/utils"
)

// handler to calculate points for a given receipt
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

// handler to check the receipt for any errors before processing
func CheckIfReceiptIsInvalid(receipt models.Receipt) []error {
	var errors []error

	if err := utils.ValidateDate(receipt.PurchaseDate); err != nil {
		errors = append(errors, err)
	}

	if err := utils.ValidateTime(receipt.PurchaseTime); err != nil {
		errors = append(errors, err)
	}

	if err := utils.ValidateAmount(receipt.Total, "total amount"); err != nil {
		errors = append(errors, err)
	}

	for i, item := range receipt.Items {
		if err := utils.ValidateItem(item, i); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}
