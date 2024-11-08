package handlers

import (
	"fetch-receipt-processor/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePointsSimple(t *testing.T) {

	var items = []models.Item{
		{
			ShortDescription: "Pepsi - 12-oz",
			Price:            "1.25",
		},
	}

	var receipt1 = models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "13:13",
		Total:        "1.25",
		Items:        items,
	}

	pointsCalculator := NewDefaultPointsCalculator()
	actualResponse := pointsCalculator.CalculatePoints(receipt1)
	expectedResponse := int64(31)

	assert.Equal(t, expectedResponse, actualResponse)

}

func TestCalculatePointsMultipleItems(t *testing.T) {

	var items = []models.Item{
		{
			ShortDescription: "Pepsi - 12-oz",
			Price:            "1.25",
		},
		{
			ShortDescription: "Dasani",
			Price:            "1.40",
		},
	}

	var receipt2 = models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items:        items,
	}

	pointsCalculator := NewDefaultPointsCalculator()
	actualResponse := pointsCalculator.CalculatePoints(receipt2)
	expectedResponse := int64(15)

	assert.Equal(t, expectedResponse, actualResponse)

}
