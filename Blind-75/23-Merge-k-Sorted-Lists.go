package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
	// Patterns:
	// MergeSort, TwoPtrs kinda, Dummy Node
	if len(lists) < 1 {
		return nil
	}

	for len(lists) > 1 {
		procLsts := []*ListNode{}

		for i := 0; i < len(lists); i += 2 {
			var l1, l2 *ListNode
			if i < len(lists) && lists[i] != nil {
				l1 = lists[i]
			}

			if i+1 < len(lists) && lists[i+1] != nil {
				l2 = lists[i+1]
			}

			merged := mergeTwoListsIntoOne(l1, l2)
			procLsts = append(procLsts, merged)
		}
		lists = procLsts
	}

	// [[1,4,5],[1,3,4],[2,6]] -> [1,1,2,3,4,4,5,6]
	// Convert "slice" to linked list
	return lists[0]
}

func mergeTwoListsIntoOne(l1, l2 *ListNode) *ListNode {
	head := &ListNode{}
	curr := head
	// each linked-list is sorted in ascending order.
	currL1, currL2 := l1, l2
	for currL1 != nil && currL2 != nil {
		if currL1.Val < currL2.Val {
			curr.Next = currL1
			currL1 = currL1.Next
		} else {
			curr.Next = currL2
			currL2 = currL2.Next
		}
		curr = curr.Next
	}

	for currL1 != nil {
		curr.Next = currL1
		currL1 = currL1.Next
		curr = curr.Next
	}

	for currL2 != nil {
		curr.Next = currL2
		currL2 = currL2.Next
		curr = curr.Next
	}
	return head.Next // cause its initially is a dummy node
}
