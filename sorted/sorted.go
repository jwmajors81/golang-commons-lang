package sorted

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](values ...T) T {
	minLength := values[0]

	for _, val := range values {
		if val < minLength {
			minLength = val
		}
	}

	return minLength
}

func Max[T constraints.Ordered](values ...T) T {
	maxLength := values[0]

	for _, val := range values {
		if val > maxLength {
			maxLength = val
		}
	}

	return maxLength
}
