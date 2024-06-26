package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/* 题目描述
题目描述很直接，就是让用前序遍历的方式遍历一个二叉树，没有多余的套路，不需要转换题义什么的
*/
 
// 递归解法
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	// 注意这里需要传入slice的指针，因为在切片扩容过程中会变更底层数组，recursion函数中传入的res会发生扩容，导致底层数组变更，但是preorderTraversal函数中的res并不会变更底层数组，因为res通过值传递到recursion函数中
	recursion(root, &res)

	return res
}

func recursion(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	recursion(node.Left, res)
	recursion(node.Right, res)
}

// 迭代写法
