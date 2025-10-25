package longest_substring_without_repeating_characters

// LengthOfLongestSubstring Time: O(n) RAM: O(1)
func LengthOfLongestSubstring(s string) int {
	n := len(s)
	maxLength := 0
	lastIndex := make([]int, 128)

	start := 0
	for end := start; end < n; end++ {
		currentChar := s[end]
		if lastIndex[currentChar] > start {
			start = lastIndex[currentChar]
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
		lastIndex[currentChar] = end + 1
	}

	return maxLength
}
