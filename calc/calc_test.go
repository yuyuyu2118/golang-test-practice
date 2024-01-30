package calc

import (
	"math"
	"testing"
)

// TestMax用のテーブル駆動テスト
func TestMax(t *testing.T) {
	t.Parallel()
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
				t.Errorf("%s: Max(%d, %d) = %d, want %d", tt.name, tt.x, tt.y, got, tt.want)
			}
		})
	}
}

// TestMin用のテーブル駆動テスト
func TestMin(t *testing.T) {
	t.Parallel()
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
				t.Errorf("%s: Min(%d, %d) = %d, want %d", tt.name, tt.x, tt.y, got, tt.want)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"0の場合", 0, false},
		{"1の場合", 1, false},
		{"素数の場合", 2, true},
		{"素数ではない場合", 4, false},
		{"負の数の場合", -1, false},
		{"数値が大きい素数の場合", 1000000009, true},
		{"数値が大きい素数ではない場合", 1000000008, false},
		{"数値がintの最大値の場合", math.MaxInt64, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.num); got != tt.want {
				t.Errorf("%s: IsPrime(%d) = %v, want %v", tt.name, tt.num, got, tt.want)
			}
		})
	}
}

func BenchmarkMax(b *testing.B) {
	x, y := int64(1), int64(math.MaxInt64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Max(x, y)
	}
}

func BenchmarkMin(b *testing.B) {
	x, y := int64(1), int64(math.MinInt64)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Min(x, y)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	max := 1000000009
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPrime(max)
	}
}
