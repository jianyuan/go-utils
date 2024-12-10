package maputils

// MapValues applies the function f to each value in the map m and returns a new map containing the results.
func MapValues[K comparable, V any, R any](m map[K]V, f func(V) R) map[K]R {
	result := make(map[K]R, len(m))
	for k, v := range m {
		result[k] = f(v)
	}
	return result
}
