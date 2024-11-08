package handler

import (
	"fetch-receipt-processor/pkg/models"
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

	actualResponse := calculatePoints(receipt1)
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

	var receipt1 = models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items:        items,
	}

	actualResponse := calculatePoints(receipt1)
	expectedResponse := int64(15)

	assert.Equal(t, expectedResponse, actualResponse)

}

func TestCalculatePointsStringParseError(t *testing.T) {

	var items = []models.Item{
		{
			ShortDescription: "Pepsi - 12-oz",
			Price:            "One",
		},
		{
			ShortDescription: "Dasani",
			Price:            "Two", // it wont be counted as the price is unable to parse into int64
		},
	}

	var receipt1 = models.Receipt{
		Retailer:     "Walgreens",
		PurchaseDate: "2022-01-02",
		PurchaseTime: "08:13",
		Total:        "2.65",
		Items:        items,
	}

	actualResponse := calculatePoints(receipt1)
	expectedResponse := int64(14) //amount reduced as item2 wont be calculated for rewards due to type mismatch

	assert.Equal(t, expectedResponse, actualResponse)
}
