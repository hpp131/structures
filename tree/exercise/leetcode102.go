package main

// 层序遍历

// 给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。

/*
	这个题目是所有二叉树层序遍历变种的模版题，即其他变种的二叉树层序遍历都可以套用这个模版，然后根据具体的题意稍加修改即可
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	depth := 0
	if root == nil {
		return res
	}
	var levelFunc func(node *TreeNode, depth int)
	levelFunc = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, make([]int, 0))
		}
		res[depth] = append(res[depth], node.Val)
		levelFunc(node.Left, depth+1)
		levelFunc(node.Right, depth+1)
	}
	levelFunc(root, depth)
	return res
}
