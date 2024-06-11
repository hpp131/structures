package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 题目描述
题目描述很直接，就是让用后序遍历的方式遍历一个二叉树，没有多余的套路，不需要转换题义什么的
*/

// left, right, mid
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	recursion(root, &res)
	return res
}

func recursion(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	recursion(node.Left, res)
	recursion(node.Right, res)
	*res = append(*res, node.Val)
}
