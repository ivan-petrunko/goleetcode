package merge_intervals

import "sort"

// Merge Time: O(n), RAM: O(n)
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	// sort intervals by left part
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	currentResult := intervals[0]
	result := [][]int{currentResult}
	i := 1
	for i < len(intervals) {
		interval := intervals[i]
		i++
		if currentResult[1] < interval[0] {
			result = append(result, interval)
			currentResult = interval
			continue
		}
		if interval[1] > currentResult[1] {
			currentResult[1] = interval[1]
		}
	}
	return result
}
