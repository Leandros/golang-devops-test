package provider

// Ptr creates a pointer to `val`.
func Ptr[T any](val T) *T {
	ret := new(T)
	*ret = val

	return ret
}
