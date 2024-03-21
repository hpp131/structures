package list

// 单项循环链表

type SingRingListNode struct {
	next  *SingRingListNode
	value int
}

func NewSingRingListNode() *SingRingListNode {
	return &SingRingListNode{
		next: nil,
	}
}

func CreateSingRingListN(n int) *SingRingListNode {
	head := NewSingRingListNode()
	p := head
	for i := 1; i < n; i++ {
		head.next = &SingRingListNode{value: i}
		head = head.next
	}
	// 用p代表尾部元素,并将尾部元素的next执行head
	p.next = head
	return head
}
func (s *SingRingListNode) InsertNode(){

}

func (s *SingRingListNode) DeleteNode(){
	
}
