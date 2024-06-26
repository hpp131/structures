package main

import (
	"container/list"
)

/*
迭代法遍历二叉树
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历
func preorderTraversal(root *TreeNode) []int {
	// 用于存放遍历结果的数组
	ans := []int{}

	if root == nil {
		return ans
	}
	// 用Go内置的双向链表来模拟栈
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}
	}
	return ans
}

// 中序遍历
func inorderTraversal(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}
	// 使用双向链表实现栈
	st := list.New()
	st.PushBack(root)
	cur := root
	for cur != nil || st.Len() > 0 {
		// 一直遍历到左子树的叶子节点
		if cur != nil {
			st.PushBack(cur)
			cur = cur.Left
		} else {
			// 将栈顶元素出栈（依次将所有左节点出栈）
			node := st.Remove(st.Back()).(*TreeNode)
			res = append(res, node.Val)
			cur = node.Right
		}
	}
	return res
}

// 后续遍历
func postorderTraversal(root *TreeNode) []int {
	// 用于存放遍历结果的数组
	ans := []int{}

	if root == nil {
		return ans
	}
	// 用Go内置的双向链表来模拟栈
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)

		ans = append(ans, node.Val)
		// 在前序遍历的基础上稍作修改，结果数组中元素的顺序为中右左
		if node.Left != nil {
			st.PushBack(node.Right)
		}
		if node.Right != nil {
			st.PushBack(node.Left)
		}
	}
	// 执行倒序操作后，顺序为左右中
	reverse(ans)
	return ans
}

func reverse(data []int) {
	left, right := 0, len(data)-1
	for left < right {
		data[left], data[right] = data[right], data[left]
		left++
		right--
	}
}

func test(root *TreeNode) []int {
	ans := []int{}

	if root == nil {
		return ans
	}
	// 用Go内置的双向链表来模拟栈
	st := list.New()
	st.PushBack(root)

	for st.Len() > 0 {
		node := st.Remove(st.Back()).(*TreeNode)
		if node.Right != nil {
			st.PushBack(node.Right)
		}
		if node.Left != nil {
			st.PushBack(node.Left)
		}

	}
	return nil
}

// 层序遍历，就是所谓的"广度优先遍历"

func layerIterate(root *TreeNode) []int {
	res := make([]int, 0)

	currentLevel := []*TreeNode{root}
	for len(currentLevel) > 0 {
		nextLevel := []*TreeNode{}
		for _, node := range currentLevel {
			// 添加当前层元素
			res = append(res, node.Val)
			// 存放下一层的元素
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		currentLevel = nextLevel
	}
	return res
}
