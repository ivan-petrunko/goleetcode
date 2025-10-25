package validate_ip_address

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		queryIP  string
		expected string
	}{
		{
			name:     "test 1",
			queryIP:  "",
			expected: "Neither",
		},
		{
			name:     "test 2",
			queryIP:  "172.16.254.1",
			expected: "IPv4",
		},
		{
			name:     "test 3",
			queryIP:  "0.0.0.0",
			expected: "IPv4",
		},
		{
			name:     "test 4",
			queryIP:  "2001:0db8:85a3:0:0:8A2E:0370:7334",
			expected: "IPv6",
		},
		{
			name:     "test 5",
			queryIP:  "256.256.256.256",
			expected: "Neither",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValidIPAddress(tt.queryIP)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ValidIPAddress(%v) = %v, but expected %v",
					tt.queryIP, got, tt.expected)
			}
		})
	}
}
