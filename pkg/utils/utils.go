package utils

import (
	"fetch-receipt-processor/internal/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// utility function to validate date
func ValidateDate(date string) error {
	if _, err := time.Parse("2006-01-02", date); err != nil {
		return fmt.Errorf("invalid purchase date: %s", date)
	}
	return nil
}

// utility function to validate time
func ValidateTime(timer string) error {
	if _, err := time.Parse("15:04", timer); err != nil {
		return fmt.Errorf("invalid purchase time: %s", timer)
	}
	return nil
}

// utility function to validate dollar value
func ValidateAmount(amount string, field string) error {
	f, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return fmt.Errorf("invalid %s: %s", field, amount)
	}
	if fmt.Sprintf("%.2f", f) != amount {
		return fmt.Errorf("invalid %s: %s (must have exactly two decimal places)", field, amount)
	}
	return nil
}

// utility function to validate price of an item
func ValidateItem(item models.Item, index int) error {
	return ValidateAmount(item.Price, fmt.Sprintf("price for item %d (%s)", index, item.ShortDescription))
}

// Rule: One point for every alphanumeric character in the retailer name.
func CalculateRetailerRewards(name string) int64 {
	count := int64(0)
	for _, c := range name {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			count++
		}
	}
	return count
}

// Rule: 50 points if the total is a round dollar amount with no cents.
func CalculateWholeAmountRewards(totalAmount string) int64 {
	total, _ := strconv.ParseFloat(totalAmount, 64)
	if total == math.Floor(total) {
		return 50
	}
	return 0
}

// Rule: 25 points if the total is a multiple of 0.25.
func CalculateMultipleOf25Rewards(totalAmount string) int64 {
	totalAmountParsed, _ := strconv.ParseFloat(totalAmount, 64)
	if math.Mod(totalAmountParsed, 0.25) == 0 {
		return 25
	}
	return 0
}

// Rule: 5 points for every two items on the receipt.
func CalculateDoubleItemRewards(items []models.Item) int64 {
	return int64(len(items) / 2 * 5)
}

// Rule: If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2
// and round up to the nearest integer. The result is the number of points earned.
func CalculateItemDescRewards(items []models.Item) int64 {
	var result int64
	for _, item := range items {
		if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
			priceFloat, _ := strconv.ParseFloat(item.Price, 64)
			result += int64(math.Ceil(priceFloat * 0.2))
		}
	}
	return result
}

// Rule: 6 points if the day in the purchase date is odd.
func CalculatePurchaseDateRewards(date string) int64 {
	purchaseDate, _ := time.Parse("2006-01-02", date)
	if purchaseDate.Day()%2 != 0 {
		return 6
	}
	return 0
}

// Rule: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func CalculatePurchaseTimeRewards(timer string) int64 {
	purchaseTime, _ := time.Parse("15:04", timer)
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		return 10
	}
	return 0
}
