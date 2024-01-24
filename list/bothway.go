package main

import "fmt"

// 双向链表

type BothWayNode struct {
	pre, next *BothWayNode
	value     int
}

// 等效于new(BothWayNode),返回结构体零值的指针
func NewBothWayNode() *BothWayNode {
	return &BothWayNode{nil, nil, 0}
}

// 创建一个长度为N的双向链表，返回头节点
func NewBothWayNodeN(value int) *BothWayNode {
	// 创建头节点
	head := NewBothWayNode()
	p := head
	for i:=1;i<value;i++{
		p.next = &BothWayNode{pre: p}
		p = p.next
	}
	// 尾节点的下一节点为nil
	p.next = nil
	return head
}

func (b *BothWayNode) Prev() *BothWayNode {
	if b.pre != nil {
		return b.pre
	}
	return nil
}

func (b *BothWayNode) Next() *BothWayNode {
	if b.next != nil {
		return b.next
	}
	return nil
}

func (b *BothWayNode) Insert(s *BothWayNode) {
	// 链表的最后一个节点后插入新元素
	if b.next == nil {
		b.next = s
		s.pre = b
		s.next = nil
	}else{
		s.next = b.next
		b.next = s
		s.pre = b
	}
}

// 删除当前节点
func (b *BothWayNode) Delete() {
	// 删除的是尾节点
	if b.next == nil {
		b.pre.next = nil
		return
	}
	// 删除的是头节点
	if b.pre == nil {
		b.next.pre = nil
		return
	}
	// 普通逻辑
	b.pre.next = b.next
	b.next.pre = b.pre

	

}
func main() {
	len := 5
	blist := NewBothWayNodeN(len)
	// fmt.Println(blist.next.next)
	for i:=1;i<len;i++{
		blist.next.value = i
		blist = blist.next
	}
	
	for i:=0;i<len;i++{
		fmt.Println(blist.value)
		blist = blist.pre
	}

	

}
