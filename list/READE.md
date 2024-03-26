1. 增
双向链表官方实现
```go
// insert inserts e after at, increments l.len, and returns e.
func (l *List) insert(e, at *Element) *Element {
	e.prev = at
	e.next = at.next
    
    // 此时的 e.prev 已经是 at 节点
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}
```
自己实现
```go
e.prev = at
e.next = at.next

at.next.prev = e
at.next = e
```
自己实现的这种方法对吗？或者有什么不好的地方？


虚拟节点dummy的优势体现在哪里？是在新增/删除操作的时候？