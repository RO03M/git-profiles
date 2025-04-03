package utils

func Find[T any](slice []T, callback func(item T, key int) bool) T {
	for index, item := range slice {
		if callback(item, index) {
			return item
		}
	}

	var zero T

	return zero
}
