package godata

// Ptr is a helper function that returns &object
func Ptr[T any](object T) *T {
	return &object
}
