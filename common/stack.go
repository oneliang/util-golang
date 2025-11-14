package common

import (
	"errors"
)

// Stack 泛型栈
type Stack[T any] struct {
	items []T
}

// NewStack 创建泛型栈
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

// Push 入栈
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop 出栈
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek 查看栈顶元素
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty 判断是否为空
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 栈大小
func (s *Stack[T]) Size() int {
	return len(s.items)
}
