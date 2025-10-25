package find_the_distance_value_between_two_arrays

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		arr1     []int
		arr2     []int
		d        int
		expected int
	}{
		{
			name:     "test 1",
			arr1:     []int{4, 5, 8},
			arr2:     []int{10, 9, 1, 8},
			d:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindTheDistanceValue(tt.arr1, tt.arr2, tt.d)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FindTheDistanceValue(%v, %v, %v) = %v, but expected %v",
					tt.arr1, tt.arr2, tt.d, got, tt.expected)
			}
		})
	}
}
