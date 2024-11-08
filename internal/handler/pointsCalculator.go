package handler

import (
	"fetch-receipt-processor/internal/models"
)

type PointsCalculator interface {
	CalculatePoints(receipt models.Receipt) int64
	CheckIfReceiptIsInvalid(receipt models.Receipt) []error
}
