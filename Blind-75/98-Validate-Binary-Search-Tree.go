package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	const INF = math.MaxInt

	var isValid func(*TreeNode, int, int) bool
	isValid = func(node *TreeNode, left, right int) bool {
		if node == nil {
			return true
		}

		if left >= node.Val || node.Val >= right {
			return false
		}

		return isValid(node.Left, left, node.Val) &&
			isValid(node.Right, node.Val, right)
	}
	// -inf < root < +inf
	return isValid(root, -INF, INF)
}
