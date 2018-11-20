package main

func IsPalindromePerm(input string) bool {
	counts := make(map[rune]int)

	for _, r := range input {
		counts[r]++
	}

	oddLen := len(input)%2 == 1
	seenOdd := false
	for _, v := range counts {
		// If value v is odd.
		if v%2 == 1 {
			if !seenOdd && oddLen {
				seenOdd = true
			} else {
				return false
			}
		}
	}
	return true
}
