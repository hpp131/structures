package main

// 给定一个二叉树 root ，返回其最大深度。

// 二叉树的 最大深度 是指从根节点到最远叶子节点的最长路径上的节点数。

/*
	从题义上来看的话，感觉用DFS会比较快，因为如果是BFS的话，遍历了所以内容，然后只取数组的长度，好像有点浪费
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	gdepth := 1
	// lt, rt := 0, 0
	var recurFunc func(node *TreeNode, depth int) int
	recurFunc = func(node *TreeNode, depth int) int {
		var res int
		// lt, rt := 0, 0
		// 递归的终止条件:到达叶子节点
		if node.Left == nil && node.Right == nil {
			return depth
		}
		if node.Left != nil {
			res = recurFunc(node.Left, depth+1)
		}
		if node.Right != nil {
			res = recurFunc(node.Right, depth+1)
		}
		// if lt > rt {
		// 	return rt
		// } else {
		// 	return lt
		// }
		return res
	}
	r := recurFunc(root, gdepth)
	return r
}
