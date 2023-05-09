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
