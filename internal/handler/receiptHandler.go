package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type processReceiptItem struct {
	Id string `json:"id"`
}

type getPointsItem struct {
	Points string `json:"points"`
}

func ProcessReceipt(c *gin.Context) {

	data := processReceiptItem{
		Id: "2000",
	}
	c.IndentedJSON(http.StatusOK, data)
}

func GetPoints(c *gin.Context) {
	data := getPointsItem{
		Points: "4000",
	}
	c.IndentedJSON(http.StatusOK, data)
}
