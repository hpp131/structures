package main

import (
	"container/list"
)

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/*
给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。

树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。
*/

type Node struct {
	Val      int
	Children []*Node
}

// 递归法
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	depth := 0
	if root == nil {
		return res
	}

	var recurFunc func(node *Node, depth int)
	recurFunc = func(node *Node, depth int) {
		if len(res) == depth {
			res = append(res, []int{})
		}
		res[depth] = append(res[depth], node.Val)
		for _, v := range node.Children {
			recurFunc(v, depth+1)
		}
	}
	recurFunc(root, depth)
	return res
}

// 迭代法
func levelOrder(root *Node) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		tmp := make([]int, 0)
		for i := 0; i < length; i++ {
			node := queue.Front().Value.(*Node)
			tmp = append(tmp, node.Val)
			// 获取到节点后马上删除
			queue.Remove(queue.Front())
			// 将下一层的节点插入到队列中
			for _, v := range node.Children {
				queue.PushBack(v)
			}
		}
		res = append(res, tmp)
	}

	return res
}



