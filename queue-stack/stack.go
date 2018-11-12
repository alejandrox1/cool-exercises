package main

import (
	"errors"
)

// Stackable is a stack interface.
type Stackable interface {
	Push(int)
	Pop() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}

// Stack is a stack of stackNodes.
type Stack struct {
	top *stackNode
}

type stackNode struct {
	value int
	prev  *stackNode
}

// Push adds a new element to the top of the stack.
func (s *Stack) Push(value int) {
	newNode := &stackNode{value: value}

	// Connect the nw node to the previous element.
	if s.top != nil {
		newNode.prev = s.top
	}

	s.top = newNode
}

// Pop removes the top of the stack.
func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return -1, errors.New("Cannot pop. Stack is empty")
	}

	ret := s.top.value
	s.top = s.top.prev

	return ret, nil
}

// Peek returns the value at the top of the stack.
func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return -1, errors.New("Cannot peek. Stack is empty")
	}

	return s.top.value, nil
}

// IsEmpty returns true if the stack is empty, otherwise it returns false.
func (s *Stack) IsEmpty() bool {
	return s.top == nil
}
