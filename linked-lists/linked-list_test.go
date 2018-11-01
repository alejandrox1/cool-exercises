package main

import (
	"reflect"
	"testing"
)

func TestLinkedListInsertAndGet(t *testing.T) {
	ll := GetLinkedList()
	vals := []int{1, 2, 3, 4, 5}
	for _, val := range vals {
		ll.Insert(val)
	}

	actual := make([]int, len(vals))
	for i := range vals {
		actual[i] = ll.Get(i)
	}

	if ll.Len() != len(vals) {
		t.Fatalf("Expected length %d. Got %d\n", len(vals), ll.Len())
	}

	if !reflect.DeepEqual(actual, vals) {
		t.Fatalf("Expected: %v. Got %v\n", actual, vals)
	}
}

func TestLinkedListConstructor(t *testing.T) {
	vals := []int{1, 2, 3, 4, 5}
	ll := GetLinkedListFromValues(vals)

	actual := make([]int, len(vals))
	for i := range vals {
		actual[i] = ll.Get(i)
	}

	if ll.Len() != len(vals) {
		t.Fatalf("Expected %d. Got %d\n", len(vals), ll.Len())
	}

	if !reflect.DeepEqual(actual, vals) {
		t.Fatalf("Expected: %v. Got %v\n", actual, vals)
	}
}

func TestRemoveNode(t *testing.T) {
	vals := []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	ll := GetLinkedListFromValues(vals)
	ll.Remove(3)

	expected := []int{1, 1, 1, 3, 3, 4, 5, 5, 6}
	actual := ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}
}
