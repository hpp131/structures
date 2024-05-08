package main

import (
	"fmt"
	"structures/skill/order/quick"
)

// 桶排序

func bucketSort(data []int)  {
	// 1. Confirm number of bucket
	buckNum := 10
	buckSlice := make([][]int, buckNum)
	// 2. insert element into buckets
	for _, element := range data {
		// Comfirm whick bucket the element is dispatched 
		buckIndex := element / 10
		buckSlice[buckIndex] = append(buckSlice[buckIndex], element)
	}
	// 3. sort each bucket
	// 这里直接使用快速排序
	for _, container := range buckSlice {
		quick.Sort(container, 0, len(container) - 1)
	}
	// 4. iterate all buckets
	i := 0 
	for _, container := range buckSlice {
		for _, element := range container {
			data[i] = element
			i++
		}
	}
}

func main()  {
	s := []int{12,34,43,85,57,90,46}
	bucketSort(s)
	fmt.Println(s)
}