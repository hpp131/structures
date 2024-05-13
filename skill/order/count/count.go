package main

import "fmt"

// Calculate prefix sum
func prefixSum(countSlice []int)  []int {
	prefixSlice := make([]int, len(countSlice))
	for i:=1;i<len(countSlice);i++ {
		if i == 1 {
			prefixSlice[i] = countSlice[i]
			continue
		}
		prefixSlice[i] = countSlice[i] + prefixSlice[i-1]
	}
	return prefixSlice
}

func countSort(data []int)  {
	// 1. iterate data and find max element in data
	maxNum := 0
	for _, element := range data {
		if element > maxNum {
			maxNum = element
		}
	}
	// 2. Make an array according to maxNum
	countSlice := make([]int, maxNum+1)
	for _, element := range data {
		countSlice[element] += 1
	}

	// 3. Calculate prefix sum
	prefixSlice := prefixSum(countSlice)
	fmt.Printf("prefixSlice is %v\n", prefixSlice)
	tmpSlice := make([]int, len(data))

	// 4. recover data's stability
	for j:=len(data)-1;j>=0;j-- {
		index := data[j]
		tmpSlice[prefixSlice[index]-1] = data[j]
		prefixSlice[index] -= 1
	}
	// 
	copy(data, tmpSlice)


}

func main()  {
	s := []int{12,34,43,85,57,90,90,46,46}
	countSort(s)
	fmt.Println(s)
}