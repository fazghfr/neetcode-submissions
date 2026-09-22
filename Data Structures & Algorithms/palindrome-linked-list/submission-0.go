/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	// o(n) to get the length of the list
	n := 0
	nTraverser := head

	for nTraverser != nil {
		nTraverser = nTraverser.Next
		n++
	}

	var middle int
	isOdd := false
	if n % 2 != 0 {
		isOdd = true
	} 

	middle = n / 2

	var pointerStack []*ListNode
	mainTraverser := head
	det := 0
	for mainTraverser != nil {
		if (isOdd && det > middle) || (!isOdd && det >= middle) {
			popped := pointerStack[len(pointerStack)-1]
			pointerStack = pointerStack[:len(pointerStack) - 1]

			if popped.Val != mainTraverser.Val {
				return false
			}
		} else {
			pointerStack = append(pointerStack, mainTraverser)
			if (isOdd && det == middle) {
				pointerStack = pointerStack[:len(pointerStack)-1]
			}
		}
		mainTraverser = mainTraverser.Next
		det++
	}
	return true
}

// odd -> middle exist
// even -> middle not exist

// odd : push to stack until middle
// after middle -> check top with current, if no match false if end of arr true


// even : push to stack until abstract middle
// do the same


