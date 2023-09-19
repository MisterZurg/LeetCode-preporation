package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// depth := 1
	var depthTraversal func(node *TreeNode, lvl int) int
	depthTraversal = func(node *TreeNode, lvl int) int {
		if node == nil {
			return lvl
		}

		return max(depthTraversal(node.Left, lvl+1), depthTraversal(node.Right, lvl+1))
	}

	return max(depthTraversal(root.Left, 1), depthTraversal(root.Right, 1))
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
