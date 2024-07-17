package main

import (
	"math"
)

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

// 正解如下:
func isValidBST(root *TreeNode) bool {
	low, high := math.MinInt, math.MaxInt
	res := recursionBST(root, low, high)
	return res
}

func recursionBST(root *TreeNode, low, high int) bool {
	// low, high = math.MinInt, math.MaxInt
	if root == nil {
		return true
	}
	// 确保root.val在开区间(low, high)内，如果不在则不是BST
	if root.Val <= low || root.Val >= high {
		return false
	}
	return recursionBST(root.Left, low, root.Val) && recursionBST(root.Right, root.Val, high)
}
