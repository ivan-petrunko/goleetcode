package valid_anagram

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		expected bool
	}{
		{
			name:     "test 1",
			s:        "anagram",
			t:        "nagaram",
			expected: true,
		},
		{
			name:     "test 2",
			s:        "rat",
			t:        "car",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsAnagram(tt.s, tt.t)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("IsAnagram(%v, %v) = %v, but expected %v",
					tt.s, tt.t, got, tt.expected)
			}
		})
	}
}
