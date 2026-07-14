package main

import "fmt"

func main() {
	// input
	nums := []int{2, 3, 5, 7, 9}

	// stage 1
	dataCh := sliceToCh(nums)

	// stage 2
	finalCh := square(dataCh)

	// stage 3
	for num := range finalCh {
		fmt.Println(num)
	}
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n // synchronous
		}
		close(out) // ends loop
	}()
	return out
}

func sliceToCh(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num // synchronous
		}
		close(out) // ends loop
	}()
	return out
}
