package main

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.Empty() {
		return zero, false
	}
	last := len(s.items) - 1
	v := s.items[last]
	s.items = s.items[:last]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	var zero T
	if s.Empty() {
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
