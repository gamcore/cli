package utils

func ForEachSlice[T any](slice []T, consumer func(T)) {
	for _, t := range slice {
		consumer(t)
	}
}

func FilterSlice[T any](slice []T, predicate func(T) bool) []T {
	var filtered []T
	for _, t := range slice {
		if predicate(t) {
			filtered = append(filtered, t)
		}
	}

	return filtered
}

func AnySlice[T any](slice []T, predicate func(T) bool) bool {
	for _, t := range slice {
		if predicate(t) {
			return true
		}
	}

	return false
}

func NoneSlice[T any](slice []T, predicate func(T) bool) bool {
	return !AnySlice[T](slice, predicate)
}

func MapSlice[T any, R any](slice []T, mapper func(T) R) []R {
	var out []R
	for _, t := range slice {
		out = append(out, mapper(t))
	}

	return out
}

func FlatMapSlice[T any, R any](slice []T, mapper func(T) []R) []R {
	var out []R
	for _, t := range slice {
		out = append(out, mapper(t)...)
	}

	return out
}
