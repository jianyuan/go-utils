// Package ptr provides a function to get a pointer to a value.
package ptr

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}
