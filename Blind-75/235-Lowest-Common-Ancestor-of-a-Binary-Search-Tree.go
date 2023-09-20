package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// All Node.val are unique.
	// It's BST meaning that
	// Parent's child
	// on the left is less or eq
	// on the right is greater or eq
	curr := root

	for {
		if curr.Val <= p.Val && curr.Val >= q.Val || curr.Val >= p.Val && curr.Val <= q.Val {
			return curr
		}

		// we have to go to left sub tree
		if curr.Val > p.Val && curr.Val > q.Val {
			curr = curr.Left
		} else if curr.Val < p.Val && curr.Val < q.Val {
			curr = curr.Right
		}
	}

	return &TreeNode{}
}
