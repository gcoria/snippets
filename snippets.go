package main

func Map[T any, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for index, current := range slice {
		result[index] = f(current)
	}
	return result
}
