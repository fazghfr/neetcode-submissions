/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func middleNode(head *ListNode) *ListNode {
	// finding len of list -> o(n)
	length := 0
	lenTraverser := head
	for lenTraverser != nil {
		length++
		lenTraverser = lenTraverser.Next
	}

	// fast and slow pointer (o(n/2))
	slow := head
	fast := head
	counter := length / 2
	if length % 2 != 0 {
		counter++
	}
	for counter != 0 {
		fast = fast.Next
		counter--
	}

	// idk, something less than o(n) times some constant that is less than n/2, so i guess worst case in general is o(n)
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}


