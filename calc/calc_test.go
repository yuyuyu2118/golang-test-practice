package calc

import (
	"testing"
)

// Test<関数名>(t *testing.T)
// この関数は、go testコマンドで実行される
// func TestMax(t *testing.T) {
// 	got := Max(1, 2)
// 	want := 2
// 	if got != want {
// 		t.Errorf("Max(1, 2) == %d, want %d", got, want)
// 	}
// }

// テーブル駆動テスト
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
				t.Errorf(("Max(%d, %d) == %d, want %d"), tt.x, tt.y, got, tt.want)
			}
		})
	}
}
