package math

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func MaxInt[T constraints.Integer](a, b T) T {
	return Max(a, b)
}

func MinInt[T constraints.Integer](a, b T) T {
	return a ^ b ^ Max(a, b)
}
