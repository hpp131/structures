package main

import "container/list"

// 基于链表实现的栈

type ListStack struct {
	// Go 内置的链表
	data *list.List
}


