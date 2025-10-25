package word_pattern

import "strings"

// WordPattern Time: O(n), RAM: O(n)
func WordPattern(pattern string, s string) bool {
	words := strings.Fields(s)

	if len(words) != len(pattern) {
		return false
	}

	charToWord := map[rune]string{}
	wordToChar := map[string]rune{}

	i := 0
	for i < len(words) {
		if char, ok := wordToChar[words[i]]; !ok {
			wordToChar[words[i]] = rune(pattern[i])
		} else if char != rune(pattern[i]) {
			return false
		}

		if word, ok := charToWord[rune(pattern[i])]; !ok {
			charToWord[rune(pattern[i])] = words[i]
		} else if word != words[i] {
			return false
		}
		i++
	}

	return true
}
