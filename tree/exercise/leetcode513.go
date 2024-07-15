package main

import "container/list"

// 给定一个二叉树的 根节点 root，请找出该二叉树的 最底层 最左边 节点的值。

// 假设二叉树中至少有一个节点。

/*
	暴力解法就是层序遍历，取层序遍历最后一层的最左侧节点
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 迭代实现的层序遍历
func findBottomLeftValue(root *TreeNode) int {
	var res [][]int
	// 先进先出队列
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := make([]int, 0)
		length := queue.Len()
		for i := 0; i < length; i++ {
			qnode := queue.Front()
			bnode := qnode.Value.(*TreeNode)
			tmp = append(tmp, bnode.Val)
			queue.Remove(qnode)
			if bnode.Left != nil {
				queue.PushBack(bnode.Left)
			}
			if bnode.Right != nil {
				queue.PushBack(bnode.Right)
			}
		}
		res = append(res, tmp)
	}
	return res[len(res)-1][0]
}
