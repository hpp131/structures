使用数组/链表实现栈结构
1. 使用数组时，数组的哪一端用作栈顶？为什么？
   -  使用数组的尾部作为栈顶，因为数组的尾部是最后一个元素，所以可以快速添加和删除元素，时间复杂度为O(1)。
2. 使用链表时（单向链表），链表的哪一端用作栈顶？为什么？
   - 使用链表的头部作为栈顶，因为链表的头部是最先添加的元素，所以可以快速添加和删除元素，时间复杂度为O(1)。
   - 如果使用链表尾部，那么在链表尾部添加/删除元素时需要先从链表头遍历到尾部然后才能执行新增/删除操作，此时时间复杂度为O(n)。
3. 如果使用双向链表来实现栈呢？