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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	levels := [][]int{}

	// BFS
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		currLevel := []int{}
		currentQueueLen := queue.Len()
		for i := 0; i < currentQueueLen; i++ {
			// Pop
			node := queue.Front()
			queue.Remove(node)

			if node.Value.(*TreeNode) != nil {
				currLevel = append(currLevel, node.Value.(*TreeNode).Val)
			}

			if node.Value.(*TreeNode) != nil && node.Value.(*TreeNode).Left != nil {
				queue.PushBack(node.Value.(*TreeNode).Left)
			}

			if node.Value.(*TreeNode) != nil && node.Value.(*TreeNode).Right != nil {
				queue.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		levels = append(levels, currLevel)
	}
	return levels
}
