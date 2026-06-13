/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    m := make(map[int]*ListNode)

	iterator := head
	listLen := -1 // 0-indexed
	for iterator != nil {
		listLen++
		m[listLen] = iterator
		iterator = iterator.Next
	}

	delIdx := listLen - n + 1
	
	if _, ok := m[delIdx-1]; !ok {
		// there is no node before target node. this is a head node deletion
		newHead := m[delIdx].Next
		return newHead 
	}
	beforeMe := m[delIdx-1]
	beforeMe.Next = beforeMe.Next.Next
	deleteMe := m[delIdx]
	deleteMe.Next = nil

	return head
}

// find the length o(n), populate map
// get the target delete idx
// delete them by direct access by idx to the map

