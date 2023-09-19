package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	curr := root

	//     []
	//   /    \
	//  []    []

	// swap
	curr.Left, curr.Right = curr.Right, curr.Left

	invertTree(curr.Left)
	invertTree(curr.Right)

	return root
}
