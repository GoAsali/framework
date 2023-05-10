package slices

type ConditionalTypes interface {
	int | string
}

func Contains[T ConditionalTypes](slice []T, target T) bool {
	for _, val := range slice {
		if target == val {
			return true
		}
	}
	return false
}

func Filter[T any](vs []T, f func(T) bool) []T {
	filtered := make([]T, 0)
	for _, v := range vs {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func Map[T any, N any](vs []T, f func(T) N) []N {
	mapped := make([]N, 0)

	for _, v := range vs {
		mapped = append(mapped, f(v))
	}

	return mapped
}

func Apply[T any](vs []T, f func(t T)) {
	for _, v := range vs {
		f(v)
	}
}
