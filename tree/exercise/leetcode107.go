package main

// 给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

/*
	相比于leetcode102，该题目仅仅多了一个数组对称元素互换的过程。二叉树遍历的代码逻辑一模一样
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
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

	// 对res数组的对称元素进行对换，i<len(res)/2写法适用于基数元素数组和偶数元素数组
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
