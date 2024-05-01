package main

func concurrentFib(n int) []int {
	s := make(chan int, n)
	go fibonacci(n, s)
	response := []int{}
	for number := range s {
		response = append(response, number)
	}

	return response
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}
