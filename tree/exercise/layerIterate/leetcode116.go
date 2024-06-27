package main

import (
	"container/list"
	// "fmt"
)

// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。

// 初始状态下，所有 next 指针都被设置为 NULL。

/*
	稍微转换一下题义，找出每个节点在树形结构中的右侧节点。根节点和最右侧节点的next都是nil，其他节点的next都不为空
*/

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	// 如果注释这个特殊情况判断，leetcode上面提交过不了，报错是"invalid memory address or nil pointer dereference"。搞不懂为啥。
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		// tmp := make([]int, 0)
		for i:=0;i<length;i++ {
			qnode := queue.Front()
			bnode := qnode.Value.(*Node)
			// bnode := queue.Remove(queue.Front()).(*Node)
			
			if i != length - 1 {
				// nil pointer dereference
				bnode.Next = qnode.Next().Value.(*Node)	
			}

			queue.Remove(qnode)
			
			// queue.Remove(queue.Front())
			if bnode.Left != nil {
				queue.PushBack(bnode.Left)
			}
			if bnode.Right != nil {
				queue.PushBack(bnode.Right)
			}
		}	
	}
	return root
}

func main() {
	root := new(Node)
	root.Left = &Node{
		Val: 1,	
	}
	root.Right = &Node{
		Val: 2,	
	}
	connect(root)
}



// func connect(root *Node) *Node {
// 	queue := list.New()
// 	queue.PushBack(root)
// 	for queue.Len() > 0 {
// 		length := queue.Len()
// 		// tmp := make([]int, 0)
// 		for i:=0;i<length;i++ {
// 			node := queue.Front()
// 			if i != length-1 && queue.Front().Next() != nil {
// 				node.Next = queue.Front().Next().Value.(*Node)
// 			}
// 			queue.Remove(queue.Front())
// 			if node.Left != nil {
// 				queue.PushBack(node.Left)
// 			}
// 			if node.Right != nil {
// 				queue.PushBack(node.Right)
// 			}
// 		}	
// 	}
// 	return root
// }
