/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	h1 := l1
	h2 := l2

	dummy := &ListNode{}
	cur := dummy
	carry := false

	var tempVal int
	for h1 != nil || h2 != nil {
		tempVal = 0

		if h1 != nil {
			tempVal += h1.Val
		}

		if h2 != nil {
			tempVal += h2.Val
		}
		if carry {
			tempVal++
		}

		if tempVal > 9 {
			tempVal -= 10
			carry = true
		} else {
			carry = false
		}
		cur.Next = &ListNode{
			Val: tempVal,
		}

		cur = cur.Next 

		if h1 != nil {
			h1 = h1.Next
		}

		if h2 != nil {
			h2 = h2.Next
		}
		
	}

	if carry {
		cur.Next = &ListNode{
			Val: 1,
		}

		cur = cur.Next 
	}

	return dummy.Next
}

// pain point :edge case handling


