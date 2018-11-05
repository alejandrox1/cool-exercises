package main

import (
	"reflect"
	"testing"
)

func TestURLifySlice(t *testing.T) {
	cases := []struct {
		input    []rune
		expected []rune
	}{
		{
			[]rune("hello my name is      "), []rune("hello%20my%20name%20is"),
		},
		{
			[]rune("hello !  "), []rune("hello%20!"),
		},
		{
			[]rune(" hello my name is        "), []rune("%20hello%20my%20name%20is"),
		},
		{
			[]rune(""), []rune(""),
		},
		{
			[]rune(" "), []rune(" "),
		},
	}

	for _, c := range cases {
		URLifySlice(c.input)
		if !reflect.DeepEqual(c.input, c.expected) {
			t.Fatalf("Expected: '%s'. Got: '%s'\n", string(c.expected), string(c.input))
		}
	}
}

func TestURLify(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"hello my name is", "hello%20my%20name%20is"},
		{"hello", "hello"},
		{"hello my name is ", "hello%20my%20name%20is%20"},
		{" hello my name is", "%20hello%20my%20name%20is"},
		{"", ""},
	}

	for _, c := range cases {
		actual := URLify(c.input)
		if actual != c.expected {
			t.Fatalf("Expected: '%s'. Got: '%s'\n", c.expected, actual)
		}
	}
}
