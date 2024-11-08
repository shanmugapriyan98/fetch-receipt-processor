package utils

import (
	"fetch-receipt-processor/internal/models"
	"testing"
)

// Generic test runner function
func runTests[T any](t *testing.T, tests []struct {
	name    string
	input   T
	wantErr bool
}, testFunc func(T) error) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := testFunc(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test failed: got error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDate(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid date", "2022-01-02", false},
		{"Invalid date format", "2022/01/02", true},
		{"Invalid date", "2022-02-30", true},
		{"Empty date", "", true},
	}

	runTests(t, tests, ValidateDate)
}

func TestValidateTime(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"Valid time", "08:13", false},
		{"Valid time", "13:01", false},
		{"Invalid time", "25:00", true},
		{"Empty time", "", true},
	}

	runTests(t, tests, ValidateTime)
}

func TestValidateAmount(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			amount string
			field  string
		}
		wantErr bool
	}{
		{"Valid amount", struct {
			amount string
			field  string
		}{"2.65", "total"}, false},
		{"Invalid format", struct {
			amount string
			field  string
		}{"2.6", "total"}, true},
		{"Non-numeric", struct {
			amount string
			field  string
		}{"abc", "total"}, true},
		{"Empty amount", struct {
			amount string
			field  string
		}{"", "total"}, true},
	}

	runTests(t, tests, func(input struct {
		amount string
		field  string
	}) error {
		return ValidateAmount(input.amount, input.field)
	})
}

func TestValidateItem(t *testing.T) {
	tests := []struct {
		name  string
		input struct {
			item  models.Item
			index int
		}
		wantErr bool
	}{
		{
			name: "Valid item",
			input: struct {
				item  models.Item
				index int
			}{models.Item{ShortDescription: "Pepsi - 12-oz", Price: "1.25"}, 0},
			wantErr: false,
		},
		{
			name: "Invalid price format",
			input: struct {
				item  models.Item
				index int
			}{models.Item{ShortDescription: "Invalid Item", Price: "1.2"}, 1},
			wantErr: true,
		},
	}

	runTests(t, tests, func(input struct {
		item  models.Item
		index int
	}) error {
		return ValidateItem(input.item, input.index)
	})
}
