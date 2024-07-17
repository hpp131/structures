package main

import "math"

// 给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

// 差值是一个正数，其数值等于两值之差的绝对值。

/*
	利用BST的属性，直接中序遍历输出一个顺序的数组，然后用双指针遍历该有序数组即可
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


// 递归进行中序遍历
func getMinimumDifference(root *TreeNode) int {
	var res []int
	recursionMinDiff(root, &res)
	min := math.MaxFloat64
	for i:=0;i<len(res)-1;i++ {
		diff := math.Abs(math.Abs(float64(res[i] - res[i+1])))
		if min > diff {
			min = diff
		}
	}
	return int(min)
}

func recursionMinDiff(node *TreeNode, res *[]int)  {
	if node.Left == nil && node.Right == nil {
		*res = append(*res, node.Val)
		return
	}
	if node.Left != nil {
		recursionMinDiff(node.Left, res)
	}
	*res = append(*res, node.Val)
	if node.Right != nil {
		recursionMinDiff(node.Right, res)
	}		
}