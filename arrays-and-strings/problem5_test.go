package main

import (
	"testing"
)

func TestIsPalindromePermTest(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"amanaplanacanalpanama", true},   // normal palindrome
		{"amanalpancaaanplanama", true},   // jumbled palindrome
		{"amanaplanacanalpanamab", false}, // not a palindrome
		{"a", true},
		{"", true},
	}

	for _, c := range cases {
		actual := IsPalindromePerm(c.input)
		if actual != c.expected {
			t.Fatalf("Input: %s. Expected: %t. Got: %t\n", c.input, c.expected, actual)
		}
	}
}
