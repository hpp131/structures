package queue
// 使用数组实现队列

import (
	"fmt"
	// "slices"
)

// 循环队列
type RingQueue struct {
	// 使用数组作为底层存储实现。也可以使用slice、linkList来实现
	arr  []int
	head int
	tail int
	size int
}

func NewRingQueue(num int) *RingQueue {
	return &RingQueue{
		arr:  make([]int, num),
		head: 0,
		tail: 0,
		size: 0,
	}
}

func (r *RingQueue) GetSize() int {
	return r.size
}

func (r *RingQueue) IsEmpty() bool {
	return r.size == 0
}

// 入队操作
func (r *RingQueue) EnQueue(elem int) {
	// 队列已满
	if r.size == len(r.arr) {
		fmt.Println("queue is full, insert failed")
	}
	if r.size < len(r.arr) {
		// 实现循环的效果
		/*
			比较取巧的实现循环的效果：tail = (tail + 1) % len(r.arr)。这种方式可以代替if的使用，能减少代码量
		*/
		if r.tail+1 > len(r.arr)-1 {
			r.arr[r.tail] = elem
			r.tail = 0
			r.size++
			return
		}
		r.arr[r.tail] = elem
		r.tail += 1
		r.size++
	}

}

// 出队操作
// 由于我们只是移动指针，并没有实际修改底层slice中的元素。因此即便是出队操作，但出队的元素仍保存在底层slice中。
func (r *RingQueue) DeQueue() (elem int) {
	// 当前队列为空，没有元素可以出队了
	if r.IsEmpty() {
		fmt.Println("queue is empty, dequeue failed")
		return
	}
	/*
		比较取巧的实现循环的效果：head= (head + 1) % len(r.arr)。这种方式可以代替if的使用，能减少代码量
	*/
	if r.head+1 > len(r.arr)-1 {
		r.head = 0
		elem = r.arr[r.head]
		r.size--
		return
	}
	r.head++
	r.size--
	return r.arr[r.head]
}

// peek队首元素
func (r *RingQueue) PeekHead() (elem int) {
	if !r.IsEmpty() {
		return r.arr[r.head]
	}
	return
}

// 获取队尾元素
func (r *RingQueue) PeekTail() (elem int) {
	if !r.IsEmpty() {
		return r.arr[r.tail]
	}
	return
}

func (r *RingQueue) DescribeQueue() []int {
	return r.arr
}
