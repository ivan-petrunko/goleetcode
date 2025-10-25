package plus_one

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		digits   []int
		expected []int
	}{
		{
			name:     "test 1",
			digits:   []int{1, 2, 3},
			expected: []int{1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PlusOne(tt.digits)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("PlusOne(%v) = %v, but expected %v",
					tt.digits, got, tt.expected)
			}
		})
	}
}
