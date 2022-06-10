package utils

func And(values ...bool) bool {
	if len(values) == 0 {
		return false
	}
	for _, val := range values {
		if !val {
			return false
		}
	}

	return true
}

func Or(values ...bool) bool {
	for _, val := range values {
		if val {
			return true
		}
	}
	return false
}
