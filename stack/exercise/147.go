package exercise

type MinStack struct {
    masterSlice []int
    slaveSlice []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        masterSlice: []int{},
        slaveSlice: []int{},
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.slaveSlice) == 0  {
        this.slaveSlice = append(this.slaveSlice, x)
        this.masterSlice = append(this.masterSlice, x)
        return
    }
    // 当x=slaveSlice的栈顶元素时也应该入栈
    if x <= this.slaveSlice[len(this.slaveSlice)-1] {
        this.slaveSlice = append(this.slaveSlice, x)
    }
    this.masterSlice = append(this.masterSlice, x)
}



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


func (this *MinStack) Top() int {
    if len(this.masterSlice) == 0 {
        return 0
    }
    return this.masterSlice[len(this.masterSlice)-1]
}


func (this *MinStack) GetMin() int {
    if len(this.slaveSlice) == 0 {
        return 0
    }
    return this.slaveSlice[len(this.slaveSlice)-1]
}
