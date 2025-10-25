package plus_one

// PlusOne Time: O(n), RAM: O(1)
func PlusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + carry
		digits[i] = sum % 10
		carry = sum / 10
	}
	if carry > 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
