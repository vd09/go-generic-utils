package generic

// MapValues extracts the values of a map into a slice.
func MapValues[K comparable, V any](m map[K]V) []V {
	values := make([]V, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

// MapKeys extracts the keys of a map into a slice.
func MapKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func SetFromSlice[K comparable](sl []K) map[K]struct{} {
	set := make(map[K]struct{})
	if len(sl) > 0 {
		for _, ele := range sl {
			set[ele] = struct{}{}
		}
	}
	return set
}

// CopySlice creates a copy of the given slice of any type.
func CopySlice[T any](src []T) []T {
	dest := make([]T, len(src))
	copy(dest, src)
	return dest
}
