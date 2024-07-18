package main

// 给你一个含重复值的二叉搜索树（BST）的根节点 root ，找出并返回 BST 中的所有 众数（即，出现频率最高的元素）。

// 如果树中有不止一个众数，可以按 任意顺序 返回。

// 假定 BST 满足如下定义：

// 结点左子树中所含节点的值 小于等于 当前节点的值
// 结点右子树中所含节点的值 大于等于 当前节点的值
// 左子树和右子树都是二叉搜索树

/*
	第一想法的思路：直接遍历这个二叉树，不管它是BST，仍然按照普通二叉树的遍历方法进行遍历（前中后都可以），然后将遍历后的数组再求众数。这个想法是可以实现的
	参考答案的思路：利用BST的性质，使用中序遍历遍历这个BST，在遍历的过程中使用两个指针对相邻的两个数进行比较（是否相等）。这样避免了第二次进行for循环。使用两个指针
	保存了上一个元素的内容，可以直接与当前元素进行比较，而不用等中序遍历完成之后再使用一个for循环来遍历，相当于是用一个指针的空间换取了一个for循环，时间成本大幅优化
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findMode(root *TreeNode) []int {
	res := make([]int, 0)
	var pre *TreeNode
	maxCount := 0
	count := 0
	var recursion func(node *TreeNode)
	// 递归法进行中序遍历
	recursion = func(node *TreeNode) {
		// 递归终止条件
		if node == nil {
			return
		}
		// 左
		recursion(node.Left)

		if pre != nil && pre.Val == node.Val {
			count++
		} else {
			// 遍历的是第一个元素
			count = 1
		}
		if count >= maxCount {
			if count > maxCount {
				res = []int{node.Val}
			} else {
				res = append(res, node.Val)
			}
			maxCount = count
		}
		pre = node
		// 右
		recursion(node.Right)
	}
	recursion(root)
	return res
}

/*
	本题的考察点：分析参考答案的解题过程，这个题目考察了使用中序遍历遍历BST从而得到一个有序的结果数组，然后使用双指针的方法来减少遍历的次数。两个指针的判断逻辑被添加到了
	中序遍历中，因此还是考察考生对二叉树遍历的理解。
*/
