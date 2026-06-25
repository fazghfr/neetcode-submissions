/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    // dividing into two linkedlists
    n := 0
    temp := head 
    for temp != nil {
        n++
        temp = temp.Next
    }

    firstBit := head
    secondBit := head
    target := n/2
    
    var prevSecondBit *ListNode
    for n != target {
        n--
        prevSecondBit = secondBit
        secondBit = secondBit.Next
    }

    prevSecondBit.Next = nil
    t2 := secondBit

    // reversing secondbit
    var prev *ListNode
    for t2 != nil {
        next := t2.Next

        t2.Next = prev
        prev = t2
        t2 = next
    }

    secondBit = prev
    t2 = secondBit
    t1 := firstBit

    dummy := &ListNode{}
    curr := dummy

    for t1 != nil && t2 != nil {
        next1 := t1.Next
        next2 := t2.Next

        curr.Next = t1
        curr = curr.Next

        curr.Next = t2
        curr = curr.Next

        t1 = next1
        t2 = next2
    }

    if t1 != nil {
        curr.Next = t1
    }
    if t2 != nil {
        curr.Next = t2
    }
    head = dummy.Next
}

// 0 1 2 3 4 5 6

// // divide into two
// 0 1 2  [] 3 4 5 6
// // reverse second bit
// 0 1 2 [] 6 5 4 3
// // merge
// 0 6 1 5 2 4 3
