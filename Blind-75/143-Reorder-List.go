package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	// slow wast pointer
	slow, fast := head, head.Next

	// Putting slow pointer on middle of ll,
	// when the faster will be on it's end
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Change the connction between right part of ll
	var prev *ListNode // := slow
	for slow != nil {
		nxt := slow.Next // Save link
		slow.Next = prev
		prev = slow
		slow = nxt
	}

	// prev end of ll
	curr := head
	for prev != nil && curr != nil {
		nxt := curr.Next
		curr.Next = prev
		nxtPrv := prev.Next
		prev.Next = nxt
		curr = nxt
		prev = nxtPrv
	}
}
