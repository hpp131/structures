package queue_test

import (
	"fmt"
	"structures/queue"
	"testing"
)

func TestRingQueue(*testing.T) {
	rq := queue.NewRingQueue(3)
	rq.EnQueue(1)
	rq.EnQueue(2)
	rq.EnQueue(3)
	rq.EnQueue(4)
	rq.DeQueue()
	rq.DeQueue()
	rq.DeQueue()
	rq.DeQueue()
	rq.DeQueue()
	fmt.Println(rq.DescribeQueue())
}
