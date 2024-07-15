package main

import "math"

// 给定一个二叉树，判断它是否是平衡二叉树
// PS:平衡二叉树 是指该树所有节点的左右子树的深度相差不超过 1。

/*
	第一思维：最大深度 - 最小深度。错误的！
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	res := getDepth(root)
	if res != -1 {
		return true
	}
	return false

}

// 获取当前节点的最大深度
func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// 获取左右子树的最大深度
	l, r := getDepth(node.Left), getDepth(node.Right)
	if l == -1 || r == -1 {
		return -1
	}
	res := math.Abs(float64(l - r))
	if res > 1 {
		return -1
	}
	return max(l, r) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
