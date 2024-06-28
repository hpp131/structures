package main

// 给定一个二叉树的 根节点 root，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。


/*
	这个题目就像解答高中的物理问题一样，需要把题义转化成实际的二叉树遍历问题。题目转换成找出每一层的最右边的节点，因此也是层序遍历的一道题目。
	我们只需要把每一层遍历结果数组的最后一个元素输出就可以了
*/


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 递归法
func rightSideView(root *TreeNode) []int {
	iterate := [][]int{}
	res := []int{}
	depth := 0
	if root == nil {
		return res
	}
	var levelFunc  func(node *TreeNode, depth int) 
	levelFunc = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(iterate) == depth {
			iterate = append(iterate, make([]int, 0))
		}
		iterate[depth] = append(iterate[depth], node.Val)
		levelFunc(node.Left, depth+1)
		levelFunc(node.Right, depth+1)
	}
	levelFunc(root, depth)
	for _, v := range iterate {
		res = append(res, v[len(v)-1])
	}
	return res
}