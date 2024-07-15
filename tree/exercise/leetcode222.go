package main

import "container/list"

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。若最底层为第 h 层，则该层包含 1~ 2h 个节点。

/*
	第一思维解法：使用层序遍历确定
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代BFS
// 未使用完全二叉树的特性，当做普通二叉树进行遍历
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// FIFO Queue
	queue := list.New()
	queue.PushBack(root)
	var res int
	// arr := make([]int, 0)
	for queue.Len() > 0 {
		res++
		qnode := queue.Front()
		bnode := qnode.Value.(*TreeNode)
		
		if bnode.Left != nil {
			queue.PushBack(bnode.Left)
			// res++
		}
		if bnode.Right != nil {
			queue.PushBack(bnode.Right)
			// res++
		}
		queue.Remove(qnode)
	}
	return res
}

// 使用完全二叉树的特性
