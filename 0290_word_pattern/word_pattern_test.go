package word_pattern

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		s        string
		expected bool
	}{
		{
			name:     "test 1",
			pattern:  "abba",
			s:        "dog cat cat dog",
			expected: true,
		},
		{
			name:     "test 2",
			pattern:  "abba",
			s:        "dog cat cat fish",
			expected: false,
		},
		{
			name:     "test 3",
			pattern:  "aaaa",
			s:        "dog cat cat dog",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WordPattern(tt.pattern, tt.s)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("WordPattern(%v, %v) = %v, but expected %v",
					tt.pattern, tt.s, got, tt.expected)
			}
		})
	}
}
