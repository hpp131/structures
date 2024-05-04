package main

import "fmt"

// []int{12,34,43,85,57,90,46}
// 选择排序：每轮循环中将最小元素方法在索引i下
func selectSort(list []int) []int {
	n := len(list)
	for i:=0;i<n-1;i++ {
		for j:=i+1;j<n;j++ {
			if list[j] < list[i] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}


func main()  {
	s := []int{12,34,43,85,57,90,46}
	fmt.Println(selectSort(s))
}