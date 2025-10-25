package valid_mountain_array

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "test 1",
			arr:      []int{1, 2, 3, 4, 5, 4, 3, 2, 1},
			expected: true,
		},
		{
			name:     "test 2",
			arr:      []int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4, 5},
			expected: false,
		},
		{
			name:     "test 3",
			arr:      []int{5, 4, 3},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidMountainArray(tt.arr)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ValidMountainArray(%v) = %v, but expected %v",
					tt.arr, got, tt.expected)
			}
		})
	}
}
