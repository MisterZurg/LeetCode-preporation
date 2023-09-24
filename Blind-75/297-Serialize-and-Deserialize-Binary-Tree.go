package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "N"
	}
	// pre-order traversal
	var sb strings.Builder

	var preorderTravetasl func(*TreeNode)
	preorderTravetasl = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("N,")
			return
		}

		// value := fmt.Sprintf("%d,", strconv.Itoa(node.Val))
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteString(",")
		preorderTravetasl(node.Left)
		preorderTravetasl(node.Right)
	}
	preorderTravetasl(root)
	// fmt.Println(sb.String())
	return sb.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodes := strings.Split(data, ",")
	if nodes[0] == "N" {
		return nil
	}

	if len(nodes) == 1 {
		nodeVal, _ := strconv.Atoi(nodes[0])
		return &TreeNode{Val: nodeVal}
	}

	idx := 0
	var preorderTravetasl func() *TreeNode
	preorderTravetasl = func() *TreeNode {
		if nodes[idx] == "N" {
			idx++
			return nil
		}

		nodeVal, _ := strconv.Atoi(nodes[idx])
		newNode := &TreeNode{Val: nodeVal}
		idx++
		newNode.Left = preorderTravetasl()
		newNode.Right = preorderTravetasl()
		return newNode
	}

	return preorderTravetasl()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
