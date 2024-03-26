package exercise

import "fmt"

// 声明节点元素
type SingleListNodes struct {
	val  int
	next *SingleListNodes
}

// 声明链表
type MyLinkedList struct {
	len       int
	dummyHead *SingleListNodes
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		dummyHead: &SingleListNodes{},
	}
}

func (this *MyLinkedList) Get(index int) int {
	// 特殊情况校验
	if index >= this.len || index < 0 {
		return -1
	}
	cur := this.dummyHead
	// 遍历到index为止，使得cur就是index位置上的元素
	for i := 0; i < index + 1; i++ {
		cur = cur.next
	}
	fmt.Println("in get")
	fmt.Println(cur.val)
	return cur.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 特殊情况校验
	if this.dummyHead.next == nil {
		this.dummyHead.next = &SingleListNodes{val: val, next: nil}
		this.len++
		fmt.Println("AddAtHead")
		this.PrintLink()
		return
	}
	// 缺少init
	this.dummyHead.next = &SingleListNodes{val: val, next: this.dummyHead.next}
	this.len++
	fmt.Println("AddAtHead")
	this.PrintLink()

}

func (this *MyLinkedList) AddAtTail(val int) {
	// 特殊情况校验
	if this.dummyHead.next == nil {
		this.dummyHead.next = &SingleListNodes{val: val, next: nil}
		this.len++
		return
	}
	cur := this.dummyHead
	for i := 0; i < this.len; i++ {
		cur = cur.next
	}
	cur.next = &SingleListNodes{val: val, next: nil}
	this.len++
	fmt.Println("AddAtTail")
	this.PrintLink()
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == this.len {
		this.AddAtTail(val)
		return
	}
	if index == 0  {
		this.AddAtHead(val)
		return
	}
	if index > this.len {
		return
	}
	cur := this.dummyHead
	// 遍历到index之前的一个元素
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	// 此时cur是index之前的一个元素
	cur.next = &SingleListNodes{val: val, next: cur.next}
	this.len++
	fmt.Println("AddAtIndex")
	this.PrintLink()
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	// 异常检测
	if index < 0 || index >= this.len {
		return
	}

	cur := this.dummyHead
	
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.next = cur.next.next
	this.len--
	fmt.Println("delete")
	this.PrintLink()
}

func (list *MyLinkedList) PrintLink(){
	cur := list.dummyHead
	for i := 0; i < list.len; i++ {
		cur = cur.next
		fmt.Println(cur.val)
	}
}


func main() {

	list := Constructor()
	list.AddAtHead(1)
	list.AddAtTail(3)
	list.AddAtIndex(1, 2)
	list.Get(1);
	list.DeleteAtIndex(1);
	list.Get(1);
	// cur := list.dummyHead
	// // for i := 0; i < list.len; i++ {
	// // 	cur = cur.next
	// // 	fmt.Println(cur.val)
	// // }

}