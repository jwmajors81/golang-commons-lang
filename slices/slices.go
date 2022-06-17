package slices

// Return a new slice with a new value appeneded
func Add[T any](original []T, value T) []T {
	dest := make([]T, len(original))
	copy(dest, original)
	dest = append(dest, value)

	return dest
}
