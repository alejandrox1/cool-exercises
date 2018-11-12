package main

import (
	"errors"
)

// Queueable represents a queue interface.
type Queueable interface {
	Add(int)
	Remove() (int, error)
	Peek() (int, error)
	IsEmpty() bool
}

// QueueInt is queue of queueNodes.
type QueueInt struct {
	head *queueNode
	tail *queueNode
}

type queueNode struct {
	value int
	prev  *queueNode
	next  *queueNode
}

// Add adds a value to the queue.
func (s *QueueInt) Add(value int) {
	newNode := &queueNode{value: value}

	// If there is a valid tail, then connect the node to it.
	if s.tail != nil {
		newNode.prev = s.tail
		s.tail.next = newNode
	}

	s.tail = newNode
	// If this is the first node in the queue, then make it the head.
	if s.head == nil {
		s.head = newNode
	}
}

// Remove pops an element from the front of the queue.
func (s *QueueInt) Remove() (int, error) {
	if s.head == nil {
		return -1, errors.New("Cannot pop. Queue is empty")
	}

	ret := s.head.value
	s.head = s.head.next
	if s.head != nil {
		s.head.prev = nil
	}

	return ret, nil
}

// Peek returns the first element in the queue.
func (s *QueueInt) Peek() (int, error) {
	if s.head == nil {
		return -1, errors.New("Cannot peek. Queue is empty")
	}

	return s.head.value, nil
}

// IsEmpty returns true if queue is empty, else it returns false.
func (s *QueueInt) IsEmpty() bool {
	return s.head == nil
}
