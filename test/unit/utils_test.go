package unit

import (
	"github.com/Folombas/modern-go-app-structure/pkg/utils"
	"testing"
)

func TestGenerateID(t *testing.T) {
	tests := []struct {
		prefix string
	}{
		{"user"},
		{"order"},
		{"pay"},
	}

	for _, tt := range tests {
		id := utils.GenerateID(tt.prefix)
		if len(id) != len(tt.prefix)+9 { // префикс- + 8 символов
			t.Errorf("GenerateID(%q) length = %d, want %d", tt.prefix, len(id), len(tt.prefix)+9)
		}
	}
}

func TestCalculateTotal(t *testing.T) {
	tests := []struct {
		amount int
		fee    int
		want   int
	}{
		{1000, 10, 1100},    // 10% комиссия
		{500, 0, 500},       // без комиссии
		{2000, 100, 4000},   // 100% комиссия
		{0, 10, 0},          // нулевая сумма
	}

	for _, tt := range tests {
		got := utils.CalculateTotal(tt.amount, tt.fee)
		if got != tt.want {
			t.Errorf("CalculateTotal(%d, %d) = %d, want %d", tt.amount, tt.fee, got, tt.want)
		}
	}
}