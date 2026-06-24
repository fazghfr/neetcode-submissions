/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
    myMap := make(map[*ListNode]bool)

	iterator := head

	for iterator != nil {
		if _, ok := myMap[iterator]; ok {
			return true
		} else {
			myMap[iterator] = true
		}
		iterator = iterator.Next
	}

	return false
}
