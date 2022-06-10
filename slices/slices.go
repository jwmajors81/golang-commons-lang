package slices

func Add[T any](original []T, value T) []T {
	dest := make([]T, len(original))
	copy(dest, original)
	dest = append(dest, value)

	return dest
}
