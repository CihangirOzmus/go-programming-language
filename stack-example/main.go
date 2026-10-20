package main

import "fmt"

func main() {
	// 1. plain stack
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

	// 2. stack as data structure
	s := &Stack[int]{}
	s.Push(11)
	s.Push(22)
	s.Push(33)

	v, _ := s.Pop()
	fmt.Printf("pop value: %v\n", v)

	v, _ = s.Peek()
	fmt.Printf("peek value: %v\n", v)
}
