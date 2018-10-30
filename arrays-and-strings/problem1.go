package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

// bruteForce will use Go's default support of unicode runes to break the
// string into individal unicode points.
func bruteForce(s string) bool {
	for i, c := range s {
		for j, cc := range s {
			if j > i && c == cc {
				return false
			}
		}
	}
	return true
}

// sortStrings uses the stdlib to sort a string.
func sortString(s string) string {
	strs := strings.Split(s, "")
	sort.Strings(strs)
	return strings.Join(strs, "")
}

// stringToRuneSlice is used to convert a string to a slice of runes. This is
// used along with the sortStringByCharacters to sort a string no matter the
// encoding.
func stringToRuneSlice(s string) []rune {
	var runes []rune
	for _, runeValue := range s {
		runes = append(runes, runeValue)
	}
	return runes
}

// sortStringByCharacters sorts a string by characters.
func sortStringByCharacters(s string) string {
	runes := stringToRuneSlice(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func isUniqueSort(s string) bool {
	s = sortStringByCharacters(s)

	for i, w := 0, 0; i < len(s); i += w {
		runValue, width := utf8.DecodeRuneInString(s[i:])
		runCompare, _ := utf8.DecodeRuneInString(s[i+width:])
		if runValue == runCompare {
			return false
		}
		w = width
	}
	return true
}

// isUniqueMap uses a map data structure to keep track of which runes have been
// seen.
func isUniqueMap(s string) bool {
	seen := make(map[rune]struct{})

	for _, r := range s {
		if _, ok := seen[r]; ok {
			return false
		}

		seen[r] = struct{}{}
	}
	return true
}

func runTests(f func(s string) bool) {
	var unique bool
	var input string

	fmt.Println("\n------------------------------------")

	input = "this is false"
	unique = f(input)
	fmt.Printf("'%s' is unique: %v\n", input, unique)

	input = "日本語"
	unique = f(input)
	fmt.Printf("'%s' is unique: %v\n", input, unique)

	input = " "
	unique = f(input)
	fmt.Printf("'%s' is unique: %v\n", input, unique)

	input = ""
	unique = f(input)
	fmt.Printf("'%s' is unique: %v\n", input, unique)

	//fmt.Println("------------------------------------\n")
}

func main() {
	runTests(bruteForce)

	runTests(isUniqueSort)

	runTests(isUniqueMap)
}
