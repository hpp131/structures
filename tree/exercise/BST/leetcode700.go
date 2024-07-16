package main

// 给定二叉搜索树（BST）的根节点 root 和一个整数值 val。

// 你需要在 BST 中找到节点值等于 val 的节点。 返回以该节点为根的子树。 如果节点不存在，则返回 null 。


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	var node *TreeNode
	// 递归返回条件
	if root.Val == val {
		return root
	}
	if root.Val > val {
		node = searchBST(root.Left, val)
	}else {
		node = searchBST(root.Right, val)
	}
	return node
}