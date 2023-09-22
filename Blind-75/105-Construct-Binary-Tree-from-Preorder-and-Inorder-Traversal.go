package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	getIndexFromInorder := func(val int) int {
		for i := range inorder {
			if val == inorder[i] {
				return i
			}
		}
		return -1
	}
	/*
	   Inorder(tree)
	       Traverse the left subtree, i.e., call Inorder(left->subtree)
	       Visit the root.
	       Traverse the right subtree, i.e., call Inorder(right->subtree)


	   Preorder(tree)
	       Visit the root.
	       Traverse the left subtree, i.e., call Preorder(left->subtree)
	       Traverse the right subtree, i.e., call Preorder(right->subtree)
	*/

	// First in preorder is always root
	root := &TreeNode{Val: preorder[0]}
	mid := getIndexFromInorder(preorder[0])
	// Once found everything before mid is in the left subtree in Inorder
	//                       after is in the right subtree

	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	return root
}
