package merge_intervals

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{
			name:      "test 1",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected:  [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "test 2",
			intervals: [][]int{{1, 4}, {4, 5}},
			expected:  [][]int{{1, 5}},
		},
		{
			name:      "test 3",
			intervals: [][]int{{1, 4}, {0, 5}},
			expected:  [][]int{{0, 5}},
		},
		{
			name:      "test 4",
			intervals: [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			expected:  [][]int{{1, 10}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Merge(%v) = %v, but expected %v",
					tt.intervals, got, tt.expected)
			}
		})
	}
}
