package main

import (
	"reflect"
	"testing"
)

func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	diff := make(map[int]int, len(a))
	for _, i := range a {
		diff[i]++
	}

	for _, i := range b {
		// If an element of b is not found.
		if _, ok := diff[i]; !ok {
			return false
		}

		diff[i]--
		if diff[i] == 0 {
			delete(diff, i)
		}
	}

	if len(diff) == 0 {
		return true
	}
	return false
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		values   []int
		target   int
		expected []int
	}{
		{[]int{1, 2, 3, 9}, 8, []int{-1, -1}},
		{[]int{1, 2, 4, 4}, 8, []int{2, 3}},
		{[]int{1}, 8, []int{-1, -1}},
	}

	for _, c := range cases {
		i, j := twoSum(c.values, c.target)
		actual := []int{i, j}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Fatalf("Expected: %v. Got %v\n", c.expected, actual)
		}
	}
}

func TestTwoSumUnsorted(t *testing.T) {
	cases := []struct {
		values   []int
		target   int
		expected []int
	}{
		{[]int{1, 2, 3, 9}, 8, []int{-1, -1}},
		{[]int{1, 2, 4, 4}, 8, []int{2, 3}},
		{[]int{1}, 8, []int{-1, -1}},
	}

	for _, c := range cases {
		i, j := twoSumUnsorted(c.values, c.target)
		actual := []int{i, j}

		if !areSlicesEqual(actual, c.expected) {
			t.Fatalf("Expected: %v. Got %v\n", c.expected, actual)
		}
	}
}
