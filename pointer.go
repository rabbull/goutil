package goutil

func Ptr[T any](val T) *T {
	return &val
}
