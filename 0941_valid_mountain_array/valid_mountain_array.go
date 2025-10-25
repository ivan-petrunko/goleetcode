package valid_mountain_array

// ValidMountainArray Time: O(n), RAM: O(1)
func ValidMountainArray(arr []int) bool {
	n := len(arr)
	if n < 3 {
		return false
	}
	isGoingUp, peakIndex := true, -1
	prev := arr[0]
	var curr int
	for i := 1; i < n; i++ {
		curr = arr[i]
		if prev == curr {
			return false
		} else if !isGoingUp && prev < curr {
			return false
		} else if prev > curr {
			isGoingUp = false
			if peakIndex == -1 {
				peakIndex = i - 1
			}
		}
		prev = curr
	}
	return peakIndex > 0 && peakIndex < n-1
}
