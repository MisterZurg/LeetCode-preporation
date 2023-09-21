package main

import "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	// RecursiveKthSmallest
	// return RecursiveKthSmallest(root, k)
	// Iterative
	return IterativeKthSmallest(root, k)
}

// Iterative
func IterativeKthSmallest(root *TreeNode, k int) int {
	stack := list.New()
	procNodes := 0

	curr := root

	for curr != nil || stack.Len() > 0 {
		// Wisit left subtree
		for curr != nil {
			stack.PushBack(curr)
			curr = curr.Left
		}
		// Pop
		poped := stack.Back()
		curr = poped.Value.(*TreeNode)
		stack.Remove(poped)
		procNodes++

		if procNodes == k {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}

// Recursive solution
func RecursiveKthSmallest(root *TreeNode, k int) int {
	inorderValues := []int{}
	// inorder traversal filling the result slice
	var inOrderTravisScott func(*TreeNode)
	inOrderTravisScott = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrderTravisScott(node.Left)
		inorderValues = append(inorderValues, node.Val)
		inOrderTravisScott(node.Right)
	}
	inOrderTravisScott(root)
	return inorderValues[k-1]
}
