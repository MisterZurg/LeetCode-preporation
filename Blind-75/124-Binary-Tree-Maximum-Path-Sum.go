package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	// Initially ... we do some more NeetCode today
	path := root.Val

	var DFS func(*TreeNode) int
	DFS = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftMax := DFS(node.Left)
		rightMax := DFS(node.Right)
		// If node value was negative, we have to update it value

		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		// if we allowed to split
		path = max(path, node.Val+leftMax+rightMax)

		return node.Val + max(leftMax, rightMax)
	}
	_ = DFS(root)

	return path
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
