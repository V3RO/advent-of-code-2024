package utils

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Abs[T Number](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}
