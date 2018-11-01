package main

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	vals := []int{}
	ll := GetLinkedListFromValues(vals)
	ll.RemoveDuplicates()
	expected := []int{}
	actual := ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 1}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicates()
	expected = []int{1}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 3}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicates()
	expected = []int{1, 3}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicates()
	expected = []int{1, 2, 3, 4, 5, 6}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}
}

func TestRemoveDuplicatesBrute(t *testing.T) {
	vals := []int{}
	ll := GetLinkedListFromValues(vals)
	ll.RemoveDuplicatesBrute()
	expected := []int{}
	actual := ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 1}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicatesBrute()
	expected = []int{1}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 3}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicatesBrute()
	expected = []int{1, 3}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}

	vals = []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	ll = GetLinkedListFromValues(vals)
	ll.RemoveDuplicatesBrute()
	expected = []int{1, 2, 3, 4, 5, 6}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %v. Got: %v\n", expected, actual)
	}
}
