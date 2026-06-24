/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var stack []*ListNode

	iterator := head
	for iterator != nil {
		stack = append(stack, iterator)
		iterator = iterator.Next
	}
	if len(stack) - 1 < 0 {
		return nil
	}
	head = stack[len(stack)-1]
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(stack) - 1 < 0 {
			cur.Next = nil
		} else {
			cur.Next = stack[len(stack) - 1]
		}
	}
	
	return head
}

