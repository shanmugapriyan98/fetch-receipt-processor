package handler

import "fmt"

// factory function to provide points calculator according to input strategy
func NewPointsCalculatorFactory(strategyType string) (PointsCalculator, error) {
	switch strategyType {
	case "one":
		return &DefaultPointsCalculator{}, nil
	default:
		return nil, fmt.Errorf("unknown strategy type: %s", strategyType)
	}
}
