package exercise

import "fmt"

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

func (r *RingQueue) getSize() int {
	return r.size
}

func (r *RingQueue) isEmpty() bool {
	return r.size == 0
}

// 入队操作
func (r *RingQueue) enqueue(elem int) {
	// 队列已满
	if r.head == r.tail && r.head != 0 {
		fmt.Println("queue is full, insert failed")
		return
	}
	if r.size < len(r.arr) {
		// 实现循环的效果
		if r.tail + 1 > len(r.arr) - 1  {
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
func (r *RingQueue) dequeue() (elem int) {
	// 当前队列为空，没有元素可以出队了
	if r.isEmpty() {
		fmt.Println("queue is empty, dequeue failed")
		return
	}
	if r.head + 1 > len(r.arr) - 1 {
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
func (r *RingQueue) peekHead() (elem int) {
	if !r.isEmpty() {
		return r.arr[r.head]
	}
	return
}

// 获取队尾元素
func (r *RingQueue) peekTail() (elem int) {
	if !r.isEmpty() {
		return r.arr[r.tail]
	}
	return
}


