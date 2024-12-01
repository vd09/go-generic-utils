package generic

type NumericAndString interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~string
}

// Min returns the minimum of two values of an ordered type.
func Min[T NumericAndString](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Min returns the minimum of two values of an ordered type.
func Max[T NumericAndString](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Helper function to check if an integer is in a slice
func ContainsSlice[T NumericAndString](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
