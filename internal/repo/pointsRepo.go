package repo

// structure definition for ReceipsAndPointsRepo repository
type PointsRepo struct {
	points map[string]int64
}

// function to return a new RecieptsAndPointsRepo instance
func NewPointsMap() *PointsRepo {
	return &PointsRepo{
		points: make(map[string]int64),
	}
}

// function to store points linked with an id
func (r *PointsRepo) StorePoints(id string, points int64) {
	r.points[id] = points
}

// function to get points linked with an id in the map
func (r *PointsRepo) GetPoints(id string) (int64, bool) {
	points, found := r.points[id]
	return points, found
}
