package string_compression

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		chars    []byte
		expected int
	}{
		{
			name:     "test 1",
			chars:    []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compress(tt.chars)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Compress(%v) = %v, but expected %v",
					tt.chars, got, tt.expected)
			}
		})
	}
}
