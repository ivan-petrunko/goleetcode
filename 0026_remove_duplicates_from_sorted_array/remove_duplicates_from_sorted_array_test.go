package remove_duplicates_from_sorted_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "test 1",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicates(tt.nums)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("RemoveDuplicates(%v) = %v, but expected %v",
					tt.nums, got, tt.expected)
			}
		})
	}
}
