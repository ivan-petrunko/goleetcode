package remove_duplicates_from_sorted_array

// RemoveDuplicates Time: O(n), RAM: O(1)
func RemoveDuplicates(nums []int) int {
	index := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
