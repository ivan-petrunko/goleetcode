package reverse_linked_list

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	tests := []struct {
		name     string
		head     *ListNode
		expected *ListNode
	}{
		{
			name:     "test 1",
			head:     LinkedListFromArray([]int{}),
			expected: LinkedListFromArray([]int{}),
		},
		{
			name:     "test 2",
			head:     LinkedListFromArray([]int{1}),
			expected: LinkedListFromArray([]int{1}),
		},
		{
			name:     "test 3",
			head:     LinkedListFromArray([]int{1, 2, 3, 4, 5}),
			expected: LinkedListFromArray([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseList(tt.head)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ReverseList(%v) = %v, but expected %v",
					tt.head, got, tt.expected)
			}
		})
	}
}
