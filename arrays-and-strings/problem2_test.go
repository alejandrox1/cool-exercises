package main

import (
	"testing"
)

func TestIsPermutation(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{"abcd", "bcda", true},
		{"abcd", "abdc", true},
		{"abcd", "aabc", false},
		{" ", " ", true},
		{"", "", true},
	}

	for _, c := range cases {
		actual := isPermutation(c.input1, c.input2)
		if actual != c.expected {
			t.Fatalf("Inputs '%s', '%s'. Expected: %t. Got %t\n",
				c.input1, c.input2, c.expected, actual)
		}
	}
}
