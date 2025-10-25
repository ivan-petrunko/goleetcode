package find_closest

import "sort"

// FindClosest Time: O(max(log(n),k)), RAM: O(1)
func FindClosest(arr []int, k int, x int) []int {
	n := len(arr)
	index := sort.SearchInts(arr, x)
	if k >= n {
		return arr
	}
	left, right := index-1, index
	for k > 0 {
		k--
		if right == n || left >= 0 && (x-arr[left]) <= (arr[right]-x) {
			left--
		} else {
			right++
		}
	}
	return arr[left+1 : right]
}
