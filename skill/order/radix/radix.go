package main

import (
	"fmt"
	"math"

)

// 以给十进制数排序为例，说明基数排序的步骤
// Get the k th num of a number
func getLastNum(num, exp int) int {
	return (num / exp) % 10
}

// radixSort depends on countSort
func couterSorting(data []int, exp int)  {
	// 存储第k个数字，因为是给10进制数排序，所以创建一个lenth=10的slice
	s := make([]int, 10)
	for _, v := range data {
		kNum := getLastNum(v, exp)
		s[kNum]++
	}
	// 计算前缀和
	for i:=1;i<len(s);i++ {
		s[i] += s[i-1]
	}
	// 准备一个临时数组存放有序的数据
	tmpSlice := make([]int, len(data))
	for j:=len(data)-1;j>=0;j-- {
		lNum := getLastNum(data[j], exp)
		tmpSlice[s[lNum]-1] = data[j]
		s[lNum]--
	}
	// 将tmpSlice中的数据还原到data中
	copy(data, tmpSlice)
}

func radixSort(data []int)  {
	maxNum := math.MinInt
	for _, v := range data {
		if v > maxNum {
			maxNum = v
		}
	}
	for exp:=1;maxNum>=exp;exp*=10 {
		couterSorting(data, exp)
	}
	
}

func main()  {
	data := []int{201700401123, 201700421145, 201701423147, 201700023147,  201800023147}
	radixSort(data)
	fmt.Println(data)
}