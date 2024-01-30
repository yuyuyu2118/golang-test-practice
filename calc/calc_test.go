package calc

import (
	"testing"
)

// TestMax用のテーブル駆動テスト
func TestMax(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"xが大きい", 2, 1, 2},
		{"yが大きい", 1, 3, 3},
		{"xとyが同じ", 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.x, tt.y); got != tt.want {
				t.Errorf("Max(%d, %d) == %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

// TestMin用のテーブル駆動テスト
func TestMin(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"xが小さい", 2, 1, 1},
		{"yが小さい", 1, 3, 1},
		{"xとyが同じ", 5, 5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.x, tt.y); got != tt.want {
				t.Errorf("Min(%d, %d) == %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
