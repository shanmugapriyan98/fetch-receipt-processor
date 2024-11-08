package handler

import (
	"fetch-receipt-processor/internal/models"
	"fetch-receipt-processor/internal/repo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// structure definition for storing result of process receipt handler
type processReceiptItem struct {
	Id string `json:"id"`
}

// structure definition for storing result of get points handler
type getPointsItem struct {
	Points int64 `json:"points"`
}

// structure definition for Receipt Handler
type ReceiptHandler struct {
	repository       repo.ReceiptsAndPointsRepo
	pointsCalculator PointsCalculator
}

// function to create an instance of Receipt Handler
func NewReceiptHandler(repo repo.ReceiptsAndPointsRepo, pointsCalculator PointsCalculator) *ReceiptHandler {
	return &ReceiptHandler{repository: repo, pointsCalculator: pointsCalculator}
}

// handler for process receipts and storing in memory
func (r *ReceiptHandler) ProcessReceipt(c *gin.Context) {
	var receipt models.Receipt

	u, err := uuid.NewV7()

	if err != nil {
		fmt.Println("Failed to generate UUID:", err)
		return
	}

	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":   "The receipt is invalid",
			"error": err,
		})
		return
	}

	uuid := u.String()

	receiptErrors := r.pointsCalculator.CheckIfReceiptIsInvalid(receipt)

	if len(receiptErrors) > 0 {
		var errorMessages []string
		for _, err := range receiptErrors {
			errorMessages = append(errorMessages, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "The receipt is invalid",
			"errors": errorMessages,
		})
		return
	}

	r.repository.StoreReceipt(uuid, receipt)

	points := r.pointsCalculator.CalculatePoints(receipt)
	r.repository.StorePoints(uuid, points)

	result := processReceiptItem{
		Id: uuid,
	}
	c.IndentedJSON(http.StatusOK, result)
}

// handler for receiving points by ID
func (r *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")
	value, found := r.repository.GetPoints(id)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "No receipt found for that id",
		})
		return
	}
	result := getPointsItem{
		Points: value,
	}
	c.IndentedJSON(http.StatusOK, result)
}
