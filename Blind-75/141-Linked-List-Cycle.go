package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	// Corner case
	if head == nil {
		return false
	}
	// two ptrs fas, slow
	slw, fst := head, head.Next

	for fst != nil && fst.Next != nil {
		if slw == fst {
			return true
		}
		slw = slw.Next
		fst = fst.Next.Next
	}
	return false
}
