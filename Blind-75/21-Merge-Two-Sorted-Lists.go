package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedHead := &ListNode{}
	curr := mergedHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	for list1 != nil {
		curr.Next = list1
		list1 = list1.Next
		curr = curr.Next
	}

	for list2 != nil {
		curr.Next = list2
		list2 = list2.Next
		curr = curr.Next
	}
	return mergedHead.Next
}
