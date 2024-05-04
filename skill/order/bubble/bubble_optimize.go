package main

import "fmt"

// []int{12,34,43,85,57,90,46}
// 该算法对原生的冒泡进行了优化，即添加了hasSwitched变量来记录某轮比较中是否发生了值交换，如果没有发生值交换，则说明排序已经完成，可以提前结束排序。

func bubbleSort(list []int) []int {
	n := len(list)
	if n <= 1 {
		return list
	}
	var hasSwitched bool
	// 控制冒泡轮数
	for i:=n-1;i>0;i-- {
		// 控制每轮中的比较次数
		for j:=0;j<i;j++ {
			if list[j] > list[j+1] {
				// swith
				list[j],list[j+1] = list[j+1],list[j]
				hasSwitched = true
			}
		}
		if	!hasSwitched {
			break
		}
	}
	return list
}

func main()  {
	s := []int{12,34,43,85,57,90,46}
	fmt.Println(bubbleSort(s))
}