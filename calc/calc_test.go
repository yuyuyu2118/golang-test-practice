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
		{"xがyより大きい正の数", 2, 1, 2},
		{"yがxより大きい正の数", 1, 3, 3},
		{"xとyが同じ正の数", 5, 5, 5},
		// 負の数の場合
		{"xがyより大きい負の数", -2, -3, -2},
		{"yがxより大きい負の数", -5, -1, -1},
		{"xとyが同じ負の数", -5, -5, -5},
		// intの最大値の場合
		{"xがyより大きく、intの最大値", math.MaxInt64, math.MaxInt64 - 1, math.MaxInt64},
		{"yがxより大きく、intの最大値", math.MaxInt64 - 1, math.MaxInt64, math.MaxInt64},
		{"xとyが同じで、intの最大値", math.MaxInt64, math.MaxInt64, math.MaxInt64},
		// intの最小値の場合
		{"xがyより大きく、intの最小値", math.MinInt64 + 1, math.MinInt64, math.MinInt64 + 1},
		{"yがxより大きく、intの最小値", math.MinInt64, math.MinInt64 + 1, math.MinInt64 + 1},
		{"xとyが同じで、intの最小値", math.MinInt64, math.MinInt64, math.MinInt64},
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
		{"xがyより小さい正の数", 2, 3, 2},
		{"yがxより小さい正の数", 5, 1, 1},
		{"xとyが同じ正の数", 5, 5, 5},
		// 負の数の場合
		{"xがyより小さい負の数", -2, -3, -3},
		{"yがxより小さい負の数", -5, -1, -5},
		{"xとyが同じ負の数", -5, -5, -5},
		// intの最大値の場合
		{"xがyより小さく、intの最大値", math.MaxInt64 - 1, math.MaxInt64, math.MaxInt64 - 1},
		{"yがxより小さく、intの最大値", math.MaxInt64, math.MaxInt64 - 1, math.MaxInt64 - 1},
		{"xとyが同じで、intの最大値", math.MaxInt64, math.MaxInt64, math.MaxInt64},
		// intの最小値の場合
		{"xがyより小さく、intの最小値", math.MinInt64, math.MinInt64 + 1, math.MinInt64},
		{"yがxより小さく、intの最小値", math.MinInt64 + 1, math.MinInt64, math.MinInt64},
		{"xとyが同じで、intの最小値", math.MinInt64, math.MinInt64, math.MinInt64},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.x, tt.y); got != tt.want {
				t.Errorf("Min(%d, %d) = %d, want %d", tt.x, tt.y, got, tt.want)
			}
		})
	}
}
