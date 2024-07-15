package main


// 给定二叉树的根节点 root ，返回所有左叶子之和。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func sumOfLeftLeaves(root *TreeNode) int {
	res := 0
	sumRecursion(root, &res, 0)
	return res
}

func sumRecursion(node *TreeNode, sum *int, label int) {
	if node != nil {
		// 确定是左叶子节点
		if node.Left == nil && node.Right == nil && label == 1{
			*sum += node.Val
		}
		sumRecursion(node.Left, sum, 1)
		sumRecursion(node.Right, sum, 2)
	}
}