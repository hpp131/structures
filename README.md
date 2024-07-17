# note

### 经常会发生的问题：
1. 数组越界  
数组越界的情况往往会发生在切片截取/对切片取索引的时候，下例是对切片取索引的时候，如果索引大于切片长度，就会发生越界。
```go
// 本例代码节选自leetcode 147 最小栈
func (this *MinStack) Pop()  {
    // 增加判断步骤为了避免表达式this.masterSlice[len(this.masterSlice)-1]产生数组越界
    if len(this.masterSlice) == 0 {
        return
    }
    popEle := this.masterSlice[len(this.masterSlice)-1]
    
    if len(this.slaveSlice) != 0 {
        if popEle == this.slaveSlice[len(this.slaveSlice)-1] {
            this.slaveSlice = this.slaveSlice[:len(this.slaveSlice)-1]
        }
    }
    this.masterSlice = this.masterSlice[:len(this.masterSlice)-1]
    return
}
```


### Caution
看到简单题目没有思路很正常，千万别怀疑自己智商，学习过程都是这样的，大家智商都差不多！  
算法刷题过程中可以学习到一些算数技巧/逻辑技巧，可以将这部分技巧应用到平时的coding中。另外，我们这样阶段的算法刷题更多的是对已有算法概念的应用，而不是创造算法，那是数学家干的事情我们暂时没有能力涉及到这方面。即使仅仅是对已有算法的应用，但是碰到大部分题目时还是没有办法AC，不要灰心，如果你能在第一次看到题目时就能结合相关的算法把题目AC，那可以说很有天分。但是大部分人，在面对题目的时候，即使有相关算法的概念，也还是不能AC，这是很正常的。快速的对题目进行AC需要一定的题目数量进行意识训练


### Q&A
1. 在leetcode上测试用例通过了，但是在提交的时候却在通过的用例上报错了，如下图所示，为什么呢？  
定义了全局变量。导致在第一个用例通过你的代码测试后，变量不会重置，导致第二个测试用例出错。建议尽量不要使用全局变量。