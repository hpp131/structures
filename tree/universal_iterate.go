package main

import "container/list"

// 使用统一模版的迭代法的前中后序遍历。iterate.go中的迭代法前中后序遍历，三着都是不一样的写法，容易混淆。
// 本节内容的原理可参考代码随想录：https://programmercarl.com/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E7%BB%9F%E4%B8%80%E8%BF%AD%E4%BB%A3%E6%B3%95.html#%E5%85%B6%E4%BB%96%E8%AF%AD%E8%A8%80%E7%89%88%E6%9C%AC
// 如果单纯看代码不能理解，可以实际的根据代码分析一个简易的二叉树，在草稿纸上一步步画出步骤进一步体会

/*
	具体思想：在入栈前，将每一个要输出的元素前防止一个nil元素作为标记。前中后序遍历代码中唯一的不同点就是压栈部分代码的顺序，其他大部分代码都一样！
*/

// 前序遍历
func PreUniversalIterate(root *TreeNode) []int {
	// res存放结果
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 使用链表模拟stack结构
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 相当于pop操作，执行操作后并不会删除链表中的实际元素
		node := stack.Back()
		// 执行删除操作，真正从链表中删除该元素
		stack.Remove(node)

		if node == nil {
			next := stack.Back()
			stack.Remove(next)
			// 实际上二叉树中的所有节点都是通过这一步输出到res数组中
			res = append(res, next.Value.(*TreeNode).Val)
			continue
		}
		// 将链表元素转换成二叉树节点
		treeNode := node.Value.(*TreeNode)
		// 按照右左中的顺序入栈，因为对于中序遍历来说，出栈时需要按照中左右的顺序出栈
		if treeNode.Right != nil {
			stack.PushBack(treeNode.Right)
		}
		if treeNode.Left != nil {
			stack.PushBack(treeNode.Left)
		}
		stack.PushBack(treeNode)
		stack.PushBack(nil)

	}
	return res
}

// 中序遍历
func MidUniversalIterate(root *TreeNode) []int {
	// res存放结果
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 使用链表模拟stack结构
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 相当于pop操作，执行操作后并不会删除链表中的实际元素
		node := stack.Back()
		// 执行删除操作，真正从链表中删除该元素
		stack.Remove(node)

		if node == nil {
			next := stack.Back()
			stack.Remove(next)
			// 实际上二叉树中的所有节点都是通过这一步输出到res数组中
			res = append(res, next.Value.(*TreeNode).Val)
			continue
		}
		// 将链表元素转换成二叉树节点
		treeNode := node.Value.(*TreeNode)
		// 按照右中左的顺序入栈，因为对于中序遍历来说，出栈时需要按照左中右的顺序出栈
		if treeNode.Right != nil {
			stack.PushBack(treeNode.Right)
		}
		stack.PushBack(treeNode)
		stack.PushBack(nil)
		if treeNode.Left != nil {
			stack.PushBack(treeNode.Left)
		}

	}
	return res
}

// 后序遍历
func PostUniversalIterate(root *TreeNode) []int {
	// res存放结果
	res := make([]int, 0)
	if root == nil {
		return res
	}
	// 使用链表模拟stack结构
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 相当于pop操作，执行操作后并不会删除链表中的实际元素
		node := stack.Back()
		// 执行删除操作，真正从链表中删除该元素
		stack.Remove(node)

		if node == nil {
			next := stack.Back()
			stack.Remove(next)
			// 实际上二叉树中的所有节点都是通过这一步输出到res数组中
			res = append(res, next.Value.(*TreeNode).Val)
			continue
		}
		// 将链表元素转换成二叉树节点
		treeNode := node.Value.(*TreeNode)
		// 按照中右左的顺序入栈，因为对于中序遍历来说，出栈时需要按照左右中的顺序出栈
		stack.PushBack(treeNode)
		stack.PushBack(nil)
		if treeNode.Left != nil {
			stack.PushBack(treeNode.Left)
		}
		if treeNode.Right != nil {
			stack.PushBack(treeNode.Right)
		}

	}
	return res
}
