package sorted

import "golang.org/x/exp/constraints"

// Returns the min value for all of the items provided
func Min[T constraints.Ordered](values ...T) T {
	minLength := values[0]

	for _, val := range values {
		if val < minLength {
			minLength = val
		}
	}

	return minLength
}

// Returns the max value for all of the items provided.
func Max[T constraints.Ordered](values ...T) T {
	maxLength := values[0]

	for _, val := range values {
		if val > maxLength {
			maxLength = val
		}
	}

	return maxLength
}
