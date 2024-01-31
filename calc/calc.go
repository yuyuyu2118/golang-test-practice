package calc

import (
	"golang.org/x/exp/constraints"
)

// Max returns the larger of x or y.
func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// Compare takes two ordered values and returns the larger and the smaller one.
func Compare[T constraints.Ordered](a, b T) (larger, smaller T) {
	if a < b {
		return b, a
	}
	return a, b
}

// IsPrime determines whether a number is a prime.
func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i <= number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}
