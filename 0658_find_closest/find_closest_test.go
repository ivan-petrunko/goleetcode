package find_closest

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "test 1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "test 2",
			arr:      []int{},
			k:        3,
			x:        3,
			expected: []int{},
		},
		{
			name:     "test 3",
			arr:      []int{1, 1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindClosest(tt.arr, tt.k, tt.x)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindClosest(%v, %d, %d) = %v, but expected %v",
					tt.arr, tt.k, tt.x, got, tt.expected)
			}
		})
	}
}
