package handler

import (
	"fetch-receipt-processor/pkg/models"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// One point for every alphanumeric character in the retailer name.
func calculateRetailerRewards(name string) int64 {
	count := int64(0)
	for _, c := range name {
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			count++
		}
	}
	return count
}

// 50 points if the total is a round dollar amount with no cents.
func calculateWholeAmountRewards(totalAmount string) int64 {
	total, err := strconv.ParseFloat(totalAmount, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return int64(0)
	}
	if total == math.Floor(total) {
		return int64(50)
	}
	return int64(0)
}

// 25 points if the total is a multiple of 0.25.
func calculateMultipleOf25Rewards(totalAmount string) int64 {
	totalAmountParsed, err := strconv.ParseFloat(totalAmount, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return int64(0)
	}
	if math.Mod(totalAmountParsed, 0.25) == 0 {
		return int64(25)
	}
	return int64(0)
}

// 5 points for every two items on the receipt.
func calculateDoubleItemRewards(items []models.Item) int64 {
	len := len(items)
	totalPairs := len / 2

	return int64(totalPairs * 5)
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2
// and round up to the nearest integer. The result is the number of points earned.
func calculateItemDescRewards(items []models.Item) int64 {
	result := int64(0)
	for _, item := range items {
		trimLen := len(strings.TrimSpace(item.ShortDescription))
		if trimLen%3 == 0 {
			priceFloat, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				fmt.Println("Error:", err)
				return int64(0)
			}
			result += int64(math.Ceil(priceFloat * 0.2))
		}
	}
	return result
}

// 6 points if the day in the purchase date is odd.
func calculatePurchaseDateRewards(date string) int64 {
	purchaseDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		fmt.Println("Error:", err)
		return int64(0)
	}
	fmt.Println(purchaseDate.Day())
	if purchaseDate.Day()%2 != 0 {
		return int64(6)
	}
	return int64(0)
}

// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
func calculatePurchaseTimeRewards(timer string) int64 {
	purchaseTime, err := time.Parse("15:04", timer)
	if err != nil {
		fmt.Println("Error:", err)
		return int64(0)
	}
	startTime, _ := time.Parse("15:04", "14:00")
	endTime, _ := time.Parse("15:04", "16:00")
	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		return int64(10)
	}
	return int64(0)
}

func calculatePoints(receipt models.Receipt) int64 {
	points := int64(0)

	points += calculateRetailerRewards(receipt.Retailer)
	points += calculateWholeAmountRewards(receipt.Total)
	points += calculateMultipleOf25Rewards(receipt.Total)
	points += calculateDoubleItemRewards(receipt.Items)
	points += calculateItemDescRewards(receipt.Items)
	points += calculatePurchaseDateRewards(receipt.PurchaseDate)
	points += calculatePurchaseTimeRewards(receipt.PurchaseTime)

	return points
}
