package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// Two Ptr prev & Curr + dummy node

	dummy := &ListNode{Next: head}
	prev, curr := dummy, head

	// Change the space between nodes

	// [dummy] [1] -> [2] -> [3] -> [4] -> [5]
	//    L     R
	//    L            R
	//    L                   R

	for i := 0; i < n; i++ {
		curr = curr.Next
	}

	// Travese until R is nil
	// for deletion of n node when right ptr will be nil
	// l ptr wiil be just before
	// [dummy] [1] -> [2] -> [3] -> [4] -> [5]
	//                        L                 R
	// prev.Next = head

	for curr != nil {
		prev = prev.Next
		curr = curr.Next
	}

	// We don't care so let's remove bond

	fmt.Println(prev.Val)

	curr = prev.Next
	prev.Next = curr.Next
	curr.Next = nil

	return dummy.Next
}
