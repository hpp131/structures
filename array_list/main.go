package main

import "fmt"



type List struct{
	value int
	next *List
}

func NewList(value int) *List{
	return &List{value: value}
}

func main()  {
	fmt.Println("hello world")

}

// 初始化单向链表:先初始化链表中的各个节点，然后将各个节点连接起来
func initList()  {
		n0 := NewList(0)
		n1 := NewList(1)
		n2 := NewList(2)
		n3 := NewList(3)
		n4 := NewList(4)
		n0.next = n1
		n1.next = n2
		n2.next = n3
		n3.next = n4
		n4.next = nil
}

// 向单项链表中插入元素（n0和n1之间插入一个节点p）
func insertElem(first , elem *List)  {
	// elem在first之后插入
	second := first.next
	elem.next = second
	first.next = elem
}

// 也可使用如下方法来插入节点,区别在于传入的参数不同，参数不同导致方法内部实现不同
// func insertElem(first , second,  elem *List)  {
// 	// elem在first之后插入
// 	elem.next = second
// 	first.next = elem
// }


// 删除单向链表中的元素（删除节点p）
func deleteElem(first, elem *List)  {
	// 删除掉elem
	// second代表原始链表中elem之后的节点
	second := elem.next
	first.next = second
}

// 按照节点值查找，并返回节点索引(以查找节点值为3的节点为例)
func  Search(head *List, value int) int {
	node := head
	for i:=0; node != nil; i++ {
		if node.value == value{
			return i
		}
		node = node.next
	}
	// 找不到返回-1
	return -1
}

