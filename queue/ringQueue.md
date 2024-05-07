### 选择实现类
1. 使用数组/slice实现队列:在这种情况下，无论使用数组头作为队列头部还是使用数组尾作为队列头部，都会存在O(n)的时间复杂度。
2. 使用链表实现队列（单向链表存在与数组同样的时间复杂度的问题）


### 循环队列解决的问题
循环队列中的循环只是一个概念，底层还是基于普通的数组/slice/linkList来实现，只不过引入了两个额外的指针。以slice实现为例：我们使用slice头部作为队列头部，slice尾部作为队列的尾部。当进行出队的时候我们只需要往后移动head指针即可，入队同样的道理。