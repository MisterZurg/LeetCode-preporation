package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	var traverseTwoTrees func(*TreeNode, *TreeNode) bool
	traverseTwoTrees = func(first, second *TreeNode) bool {
		if first == nil && second == nil {
			return true
		}

		if first != nil && second == nil || first == nil && second != nil {
			return false
		}

		if first != nil && second != nil && first.Val != second.Val {
			return false
		}

		return traverseTwoTrees(first.Left, second.Left) && traverseTwoTrees(first.Right, second.Right)
	}

	return traverseTwoTrees(q, p)
}
