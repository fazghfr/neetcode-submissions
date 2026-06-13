/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // optimal solution
	// two pointer while maintaining length n between the two

	left := 0
	right := 0

	curLeft := head
	curRight := head
	for right - left != n {
		curRight = curRight.Next
		right++
	}
	if curRight == nil {
        return head.Next
    }

	for curRight.Next != nil {
		curLeft = curLeft.Next
		curRight = curRight.Next
		left++
		right++
	}

	deleteMe := curLeft.Next
	curLeft.Next = curLeft.Next.Next
	deleteMe.Next = nil
	return head
}
