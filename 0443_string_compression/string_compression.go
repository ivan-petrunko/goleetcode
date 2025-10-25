package string_compression

import "strconv"

// Compress Time: O(n), RAM: O(1)
func Compress(chars []byte) int {
	n := len(chars)
	idx := 0
	i := 0
	for i < n {
		ch := chars[i]
		count := 0
		for i < n && chars[i] == ch {
			count++
			i++
		}
		if count == 1 {
			chars[idx] = ch
			idx++
		} else {
			chars[idx] = ch
			idx++
			for _, digit := range []byte(strconv.Itoa(count)) {
				chars[idx] = digit
				idx++
			}
		}
	}
	return idx
}
