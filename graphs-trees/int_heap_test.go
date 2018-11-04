package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestIntHeap(t *testing.T) {
	h := &IntHeap{3, 2, 1}
	heap.Init(h)
	heap.Push(h, 5)

	actual := []int{}
	for h.Len() > 0 {
		actual = append(actual, heap.Pop(h).(int))
	}

	expected := []int{1, 2, 3, 5}
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}
}
