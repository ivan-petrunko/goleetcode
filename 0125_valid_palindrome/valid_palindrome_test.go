package valid_palindrome

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "test 1",
			s:        "amanaplanacanalpanama",
			expected: true,
		},
		{
			name:     "test 2",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "test 3",
			s:        "race a car",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsPalindrome(tt.s)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("IsPalindrome(%v) = %v, but expected %v",
					tt.s, got, tt.expected)
			}
		})
	}
}
