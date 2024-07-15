package main

import "math"

// 给定一个二叉树，找出其最小深度。

// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func minDepth(root *TreeNode) int {
    // 递归的终止条件
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil {
        return 1
    }
    
	minD := math.MaxInt32
    // ans := math.MaxInt32
	if root.Left != nil {
		minD = min(minDepth(root.Left), minD)		
	}
	if root.Right != nil {
		minD = min(minDepth(root.Right), minD)
	}
	
	return minD + 1
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}