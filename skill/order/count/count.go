package main

import "fmt"


func countSort(data []int)  {
	// 1. iterate data and find max element in data
	maxNum := 0
	for _, element := range data {
		if element > maxNum {
			maxNum = element
		}
	}
	// 2. make an array according to maxNum
	tmpSlice := make([]int, maxNum+1)
	for _, element := range data {
		tmpSlice[element] += 1
	}
	// 3. fill sorted element to origin data
	dataIndex := 0
	for index, value := range tmpSlice {
		for loop:=0;loop<value;loop++ {
			data[dataIndex] = index
			dataIndex++
		}
	}
	// 恢复原数组的稳定性
	/*
	
	*/
}

func main()  {
	s := []int{12,34,43,85,57,90,46,46}
	countSort(s)
	fmt.Println(s)
}