package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// 1) Traverse util we found a key with val eq = subTree root Val
	// 2) Once found check if subTreezz are eq
	return searchSubTree(root, subRoot)
}

func searchSubTree(node, subRoot *TreeNode) bool {
	if node == nil {
		return false
	}

	var flag bool
	if node.Val == subRoot.Val {
		flag = isTwoTreesEqual(node, subRoot)
	}

	return flag || searchSubTree(node.Left, subRoot) || searchSubTree(node.Right, subRoot)
}

func isTwoTreesEqual(first, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil && second != nil || first != nil && second == nil {
		return false
	}

	if first != nil && second != nil && first.Val != second.Val {
		return false
	}

	return isTwoTreesEqual(first.Left, second.Left) && isTwoTreesEqual(first.Right, second.Right)
}
