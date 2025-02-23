// Package sliceutils provides utility functions for working with slices.
package sliceutils

// Map applies the function f to each element in the slice s and returns a new slice containing the results.
func Map[T, R any](f func(T) R, s []T) []R {
	result := make([]R, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

// Find returns the first element in the slice that satisfies the predicate f.
// If no element satisfies the predicate, it returns nil.
func Find[T any](f func(T) bool, s []T) *T {
	for _, v := range s {
		if f(v) {
			return &v
		}
	}
	return nil
}
