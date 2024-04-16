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
