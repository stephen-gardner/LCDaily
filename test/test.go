package test

func CloneSlice[T any](arr []T) []T {
	dupe := make([]T, len(arr))
	copy(dupe, arr)
	return dupe
}
