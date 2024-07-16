package main

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	var res bool
	// 递归终止条件:遍历到叶子节点后就返回
	if root.Left == nil && root.Right == nil {
		res = true
	}
	// 单层递归逻辑
	if root.Left != nil {
		if root.Left.Val < root.Val {
			if isValidBST(root.Left) {
				res = true
			}
		} else {
			res = false
		}
	}
	if root.Right != nil {
		if root.Right.Val > root.Val {
			if isValidBST(root.Right) {
				res = true
			}
		} else {
			res = false
		}
	}

	return res
}

// 以上判断逻辑有问题: 当前仅仅是比较node.left.val node.val node.right.val，并没有考虑到node的父节点
