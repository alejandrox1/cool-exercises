package main

// URLifySlice assumes that there is enough space at the end of the string for the
// extra runes.
func URLifySlice(input []rune) {
	inWord := false
	slowPtr := len(input) - 1

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == rune(' ') {
			// Ignore any white space at the end. Only start replacing once you
			// have come across a word.
			if inWord {
				input[slowPtr-2] = rune('%')
				input[slowPtr-1] = rune('2')
				input[slowPtr] = rune('0')
				slowPtr -= 3
			}
		} else {
			if !inWord {
				inWord = true
			}
			input[slowPtr] = input[i]
			slowPtr--
		}
	}
}

// URLify uses O(n) extra space.
func URLify(input string) string {
	count := 0
	spaces := 0
	for _, r := range input {
		count++
		if r == rune(' ') {
			spaces++
		}
	}

	// Add +2 for each space.
	temp := make([]rune, count+(2*spaces))
	i := 0
	for _, r := range input {
		if r == rune(' ') {
			temp[i] = rune('%')
			temp[i+1] = rune('2')
			temp[i+2] = rune('0')
			i += 3
		} else {
			temp[i] = r
			i++
		}
	}

	return string(temp)
}
