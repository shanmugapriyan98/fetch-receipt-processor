package routers

import (
	"fetch-receipt-processor/internal/handlers"

	"github.com/gin-gonic/gin"
)

func InitRouter(receiptHandler *handlers.ReceiptHandler) *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", receiptHandler.ProcessReceipt)
	router.GET("/receipts/:id/points", receiptHandler.GetPoints)
	return router
}
