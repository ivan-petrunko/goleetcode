package valid_palindrome

import "strings"

// IsPalindrome Time: O(n), RAM: O(n)
func IsPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !('a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9') {
			i++
		}
		for i < j && !('a' <= s[j] && s[j] <= 'z' || '0' <= s[j] && s[j] <= '9') {
			j--
		}
		//fmt.Printf("%c vs. %c\n", s[i], s[j])
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
