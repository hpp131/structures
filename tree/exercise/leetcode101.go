package main

import (
	"container/list"
)

// 给你一个二叉树的根节点 root ， 检查它是否轴对称。


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/*
错误的解答思路:
1. 肯定是奇数个节点（除了root节点外其余层都是偶数节点）
2. 每一层遍历后的数组都是对称的

特殊示例
[1,2,2,null,3,null,3]==>[[1] [2 2] [3 3]]。在最后一层有四个元素（从左往右为null, 3, null, 3)， 因为最后两个3并不对称，但是按照以上的判断思路，结果是对称的。因此该解法的思路不正确！
*/
func isSymmetric(root *TreeNode) bool {
	s := [][]int{}
	// 层序遍历
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		tmp := make([]int, 0)
		length := queue.Len()
		for i := 0; i < length; i++ {
			qnode := queue.Front()
			bnode := qnode.Value.(*TreeNode)
			queue.Remove(qnode)
			tmp = append(tmp, bnode.Val)
			if bnode.Left != nil {
				queue.PushBack(bnode.Left)
			}
			if bnode.Right != nil {
				queue.PushBack(bnode.Right)
			}
		}
		s = append(s, tmp)
	}
	// 遍历最后的结果数组，在遍历过程中判断二叉树是否对称
	for i, v := range s {
		// 跳过root节点
		if i == 0 {
			continue
		}
		// 判断是否偶数个
		if len(v)%2 != 0 {
			return false
		} else {
			left, right := 0, len(v)-1
			for left < right {
				if v[left] != v[right] {
					return false
				} else {
					left++
					right--
				}
			}
		}
	}
	return true
}

/*
正确思路：

*/
func isSymmetric(root *TreeNode) bool {
    return check(root, root)
}

func check(p, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if p == nil || q == nil {
        return false
    }
    return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left) 
}
