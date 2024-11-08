package repo

import "fetch-receipt-processor/internal/models"

// structure definition for ReceipsAndPointsRepo repository
type ReceiptsAndPointsRepo struct {
	receipts map[string]models.Receipt
	points   map[string]int64
}

// function to return a new RecieptsAndPointsRepo instance
func NewPointsMap() *ReceiptsAndPointsRepo {
	return &ReceiptsAndPointsRepo{
		receipts: make(map[string]models.Receipt),
		points:   make(map[string]int64),
	}
}

// function to store receipt linked with an id
func (r *ReceiptsAndPointsRepo) StoreReceipt(id string, receipt models.Receipt) {
	r.receipts[id] = receipt
}

// function to store points linked with an id
func (r *ReceiptsAndPointsRepo) StorePoints(id string, points int64) {
	r.points[id] = points
}

// function to get points linked with an id in the map
func (r *ReceiptsAndPointsRepo) GetPoints(id string) (int64, bool) {
	points, found := r.points[id]
	return points, found
}
