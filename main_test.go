package main

import (
	"bytes"
	"encoding/json"
	"fetch-receipt-processor/internal/handlers"
	"fetch-receipt-processor/internal/repo"
	"fetch-receipt-processor/internal/routers"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func sendRequest(t *testing.T, r *gin.Engine, method string, url string, body []byte) (*httptest.ResponseRecorder, []byte) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	respBody, err := io.ReadAll(w.Body)
	require.NoError(t, err)

	return w, respBody
}

var (
	receiptHandler *handlers.ReceiptHandler
	router         *gin.Engine
)

func TestMain(m *testing.M) {
	pointsCalc := handlers.NewDefaultPointsCalculator()
	repo := repo.NewPointsMap()
	receiptHandler = handlers.NewReceiptHandler(*repo, pointsCalc)

	router = routers.InitRouter(receiptHandler)

	m.Run()
}

func TestProcessReceipt(t *testing.T) {

	// Test data taken from examples directory
	successRequest := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{"shortDescription": "Mountain Dew 12PK", "price": "6.49"},
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
			{"shortDescription": "Knorr Creamy Chicken", "price": "1.26"},
			{"shortDescription": "Doritos Nacho Cheese", "price": "3.35"},
			{"shortDescription": "Klarbrunn 12-PK 12 FL OZ", "price": "12.00"}
		],
		"total": "35.35"
	}`

	// Send POST Request for Process Receipt Flow
	w, respBody := sendRequest(t, router, "POST", "/receipts/process", []byte(successRequest))
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP STATUS OK for Receipt Processing")

	// Assign response body to result and validate errors
	var result handlers.ProcessReceiptItem
	err := json.Unmarshal(respBody, &result)
	require.NoError(t, err, "Unable to parse response JSON")
	require.NotEmpty(t, result.Id, "Receipt ID missing in the response JSON")

	// Send GET request for Get Points Flow
	url := fmt.Sprintf("/receipts/%s/points", result.Id)
	w, respBody = sendRequest(t, router, "GET", url, nil)
	assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP STATUS OK for Get Points")

	// Validate points and check whether we are receiving expected points
	expectedResponse := `{"points":28}`
	require.JSONEq(t, expectedResponse, string(respBody), "Mismatch in points received from Get Points endpoint")
}

func TestProcessReceipt2(t *testing.T) {

	invalidRequest := `{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
			{"shortDescription": "Mountain Dew 12PK", "price": "Six"}, //invalid price
			{"shortDescription": "Emils Cheese Pizza", "price": "12.25"},
		],
		"total": "35.35"
	}`

	w, _ := sendRequest(t, router, "POST", "/receipts/process", []byte(invalidRequest))
	assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP BAD REQUEST for Invalid Receipt")
}

func TestProcessReceipt3(t *testing.T) {
	w, _ := sendRequest(t, router, "GET", "receipts/190/points", nil)
	assert.Equal(t, http.StatusNotFound, w.Code, "Expected HTTP STATUS NOT FOUND for Get Points with ID which does not exists")

}
