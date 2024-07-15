package main

import (
	"container/list"
	// "fmt"
)

// 给你一棵二叉树的根节点 root ，翻转这棵二叉树，并返回其根节点。

/*
	题意转换：所谓翻转二叉树，就是把每个节点的左子结点与右子节点位置互换。因此到了这里，这道题目就是一道二叉树遍历的题目
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

// 迭代层序遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			qnode := queue.Front()
			bnode := qnode.Value.(*TreeNode)
			// 交换左右子节点
			bnode.Right, bnode.Left = bnode.Left, bnode.Right

			if bnode.Left != nil {
				queue.PushBack(bnode.Left)
			}
			if bnode.Right != nil {
				queue.PushBack(bnode.Right)
			}
			queue.Remove(qnode)
		}
	}
	return root
}

// 迭代前序遍历
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	
	// 中左右
	st := list.New()
	st.PushBack(root)
	for st.Len() > 0 {
		qnode := st.Back()
		st.Remove(qnode)
		if qnode.Value == nil {
			// 从栈中pop出nil的下一个元素
			next := st.Back()
			val := next.Value.(*TreeNode)
			val.Left, val.Right = val.Right, val.Left
			st.Remove(next)
			continue
		}

		bnode := qnode.Value.(*TreeNode)
		if bnode.Right != nil {
			st.PushBack(bnode.Right)
		}
		if bnode.Left != nil {
			st.PushBack(bnode.Left)
		}
		st.PushBack(bnode)
		st.PushBack(nil)
	}
	return root
}


func main() {
	root := new(TreeNode)
	root.Left = &TreeNode{
		Val: 1,	
	}
	root.Right = &TreeNode{
		Val: 2,	
	}
	invertTree1(root)
}
