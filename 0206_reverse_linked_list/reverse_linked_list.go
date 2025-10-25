package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// LinkedListFromArray Convert array into linked list
func LinkedListFromArray(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	root := &ListNode{arr[0], nil}
	curr := root
	for i := 1; i < len(arr); i++ {
		newNode := &ListNode{arr[i], nil}
		curr.Next = newNode
		curr = newNode
	}
	return root
}

// ReverseList Time: O(n), RAM: O(1)
func ReverseList(head *ListNode) *ListNode {
	var (
		tmp *ListNode
		cur = head
	)
	for cur != nil {
		next := cur.Next
		cur.Next = tmp
		tmp = cur
		cur = next
	}
	return tmp
}
