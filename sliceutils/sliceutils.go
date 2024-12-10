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
