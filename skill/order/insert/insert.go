package main

import "fmt"

// []int{12,34,43,85,57,90,46}
// 插入排序：从一个无序的数组中取一个元素并将这个元素插入到一个有序的数组中



func insertSort(list []int)  []int {
	n := len(list)
	for i:=0;i<n-1;i++ {
		if list[i+1] < list[i] {
			for k:=i;k>=0;k-- {
				if list[k] > list[k+1] {
					list[k], list[k+1] = list[k+1], list[k]
				}else {
					break
				}
			}
		}
	}
	return list
}


func main()  {
	s := []int{12,34,43,85,57,90,46,55}
	fmt.Println(insertSort(s))
}
