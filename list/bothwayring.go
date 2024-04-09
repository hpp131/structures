// 双向循环链表

package list

import "fmt"

type RingListNode struct {
	pre, next *RingListNode
	value     int
}

func NewRingListNode() *RingListNode {
	return &RingListNode{}
}

// 创建一个长度为n的循环链表
func NewRingListN(n int) *RingListNode {
	r := NewRingListNode()
	// 循环N-1次即可。整个逻辑可在纸上画出来体会
	for i := 0; i < n; i++ {
		r.next = &RingListNode{pre: r}
		r = r.next
	}
	r.next = r
	r.pre = r
	return r
}

// 添加节点
func Add(r, target *RingListNode){
	// r表示新节点，target表示r在其后面插入
	r.next = target.next
	target.next = r
	r.pre = target
}

func main() {
	fmt.Println("this is ring.go")
}
