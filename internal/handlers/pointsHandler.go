package handlers

import (
	"fetch-receipt-processor/internal/models"
	"fetch-receipt-processor/pkg/utils"
)

type DefaultPointsCalculator struct{}

func NewDefaultPointsCalculator() PointsCalculator {
	return &DefaultPointsCalculator{}
}

// handler to calculate points for a given receipt
func (p *DefaultPointsCalculator) CalculatePoints(receipt models.Receipt) int64 {
	points := int64(0)

	points += utils.CalculateRetailerRewards(receipt.Retailer)
	points += utils.CalculateWholeAmountRewards(receipt.Total)
	points += utils.CalculateMultipleOf25Rewards(receipt.Total)
	points += utils.CalculateDoubleItemRewards(receipt.Items)
	points += utils.CalculateItemDescRewards(receipt.Items)
	points += utils.CalculatePurchaseDateRewards(receipt.PurchaseDate)
	points += utils.CalculatePurchaseTimeRewards(receipt.PurchaseTime)

	return points
}

// handler to check the receipt for any errors before processing
func (p *DefaultPointsCalculator) CheckIfReceiptIsInvalid(receipt models.Receipt) []error {
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
