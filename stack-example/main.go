package main

import "fmt"

func main() {
	stack := []int{1, 2, 3}

	// push
	stack = append(stack, 10)
	stack = append(stack, 20)
	stack = append(stack, 30)
	fmt.Printf("initial stack: %v\n", stack)

	// peek
	top := stack[len(stack)-1]
	fmt.Printf("top: %v\n", top)

	// pop
	stack = stack[:len(stack)-1]
	fmt.Printf("final stack: %v\n", stack)
}
