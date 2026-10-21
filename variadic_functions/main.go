package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(1, 2, 3))

	vals := []int{4, 5, 6}
	fmt.Println(sum(vals...))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
