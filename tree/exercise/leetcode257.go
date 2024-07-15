package main

import (
	"strconv"
)

// 给你一个二叉树的根节点 root ，按 任意顺序 ，返回所有从根节点到叶子节点的路径。

/*
	Comprehession: 对每个叶子节点/非叶子节点来说，其父节点是唯一的。对于每一个父节点来说，其孩子节点不是唯一的。所以利用唯一性，我们应该从第一种
	思路入手
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var res []string

func binaryTreePaths(root *TreeNode) []string {
	res = []string{}
	recursion(root, "")
	return res
}

func recursion(node *TreeNode, path string)  {
	if node != nil {
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
		}else {
			path += "->"
			recursion(node.Left, path)
			recursion(node.Right, path)
		}
	}

}
