package handler

import (
	"fetch-receipt-processor/internal/repo"
	"fetch-receipt-processor/pkg/models"
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
	repo repo.ReceiptsAndPointsRepo
}

// function to create an instance of Receipt Handler
func NewReceiptHandler(repo repo.ReceiptsAndPointsRepo) *ReceiptHandler {
	return &ReceiptHandler{repo: repo}
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

	receiptErrors := CheckIfReceiptIsInvalid(receipt)

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

	r.repo.StoreReceipt(uuid, receipt)

	points := calculatePoints(receipt)
	r.repo.StorePoints(uuid, points)

	result := processReceiptItem{
		Id: uuid,
	}
	c.IndentedJSON(http.StatusOK, result)
}

// handler for receiving points by ID
func (r *ReceiptHandler) GetPoints(c *gin.Context) {
	id := c.Param("id")
	value, found := r.repo.GetPoints(id)
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
