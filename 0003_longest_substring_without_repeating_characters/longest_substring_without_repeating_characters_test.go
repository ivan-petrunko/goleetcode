package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected int
	}{
		{
			name:     "test 1",
			str:      "abcabcbb",
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LengthOfLongestSubstring(tt.str)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("LengthOfLongestSubstring(%v) = %v, but expected %v",
					tt.str, got, tt.expected)
			}
		})
	}
}
