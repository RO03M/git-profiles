package utils

func Filter[T any](slice []T, callback func(item T, key int) bool) []T {
	var newArray []T = make([]T, 0, len(slice))

	for index, item := range slice {
		if callback(item, index) {
			newArray = append(newArray, item)
		}
	}

	return newArray
}
