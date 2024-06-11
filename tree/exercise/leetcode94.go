package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 题目描述
题目描述很直接，就是让用中序遍历的方式遍历一个二叉树，没有多余的套路，不需要转换题义什么的
*/

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	recursion(root, &res)
	return res
}

func recursion(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	recursion(node.Left, res)
	*res = append(*res, node.Val)
	recursion(node.Right, res)
}
