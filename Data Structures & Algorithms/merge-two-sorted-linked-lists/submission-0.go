/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    head1 := list1
	head2 := list2

	dummy := &ListNode{}
	tail := dummy

	// var prev *ListNode
	// var mergeHead *ListNode
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			tail.Next = head1
			head1 = head1.Next
		} else {
			tail.Next = head2
			head2 = head2.Next
		}

		tail = tail.Next
	}

	// last cleanup
	for head1 != nil {
		tail.Next = head1
		head1 = head1.Next
		tail = tail.Next
	}
	for head2 != nil {
		tail.Next = head2
		head2 = head2.Next
		tail = tail.Next
	}

	return dummy.Next
}

// compare head1 and head2
// "append" the smallest always

