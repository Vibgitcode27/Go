package main

func getLast[T any](s []T) T {
	store := make([]T, len(s))

	if len(store) == 0 {
		var zero T

		return zero
	}

	return store[len(store)-1]
}
