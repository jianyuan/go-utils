// Package ptr provides a function to get a pointer to a value.
package ptr

// Ptr returns a pointer to v.
func Ptr[T any](v T) *T {
	return &v
}

// Value returns the value of a pointer, or the zero value if the pointer is nil.
func Value[T any](p *T) T {
	if p == nil {
		var zeroValue T
		return zeroValue
	}
	return *p
}
