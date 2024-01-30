package calc

import (
	"math"
	"testing"
)

// TestMax用のテーブル駆動テスト
func TestMax(t *testing.T) {
	tests := []struct {
		name string
		x    int64
		y    int64
		want int64
	}{
		// 正の数の場合
		{"xが大きい", 2, 1, 2},
		{"yが大きい", 1, 3, 3},
		{"xとyが同じ", 5, 5, 5},
		// 負の数の場合
		{"xが大きい", -2, -3, -2},
		{"yが大きい", -5, -1, -1},
		{"xとyが同じ", -5, -5, -5},
		// intの最大値の場合
		{"xが大きい", math.MaxInt64, 1, math.MaxInt64},
		{"yが大きい", 1, math.MaxInt64, math.MaxInt64},
		{"xとyが同じ", math.MaxInt64, math.MaxInt64, math.MaxInt64},
		// intの最小値の場合
		{"xが大きい", 1, math.MinInt64, 1},
		{"yが大きい", math.MinInt64, 1, 1},
		{"xとyが同じ", math.MinInt64, math.MinInt64, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.x, tt.y); got != tt.want {
				t.Errorf("Max(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}

// TestMin用のテーブル駆動テスト
func TestMin(t *testing.T) {
	tests := []struct {
		name string
		x    int64
		y    int64
		want int64
	}{
		// 正の数の場合
		{"xが小さい", 2, 1, 1},
		{"yが小さい", 1, 3, 1},
		{"xとyが同じ", 5, 5, 5},
		// 負の数の場合
		{"xが小さい", -2, -3, -3},
		{"yが小さい", -5, -1, -5},
		{"xとyが同じ", -5, -5, -5},
		// intの最大値の場合
		{"xが小さい", 1, math.MaxInt64, 1},
		{"yが小さい", math.MaxInt64, 1, 1},
		{"xとyが同じ", math.MaxInt64, math.MaxInt64, math.MaxInt64},
		// intの最小値の場合
		{"xが小さい", math.MinInt64, 1, math.MinInt64},
		{"yが小さい", 1, math.MinInt64, math.MinInt64},
		{"xとyが同じ", math.MinInt64, math.MinInt64, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.x, tt.y); got != tt.want {
				t.Errorf("Min(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
