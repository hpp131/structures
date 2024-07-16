package main



// 给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值
// 相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func hasPathSum(root *TreeNode, targetSum int) bool {
	res := hasRecursion(root, 0, targetSum)
	return res
}


func hasRecursion(node *TreeNode, sum int, total int) bool {
	if node != nil {
		sum += node.Val
		if hasRecursion(node.Left, sum, total) {
			return true
		}
		if hasRecursion(node.Right, sum, total) {
			return true
		}
		if node.Left == nil && node.Right == nil {
			// 遇到叶子结点直接把相加和追加到结果集中
			if sum == total {
				return true
			}
		}
	}
	return false
}