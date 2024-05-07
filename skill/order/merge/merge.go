package main

import "fmt"




func mergeSort(data []int, lo, hi int) {
	// 递归终止条件
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	// 先将左边排序好
	mergeSort(data, lo, mid)
	// 再将右边排序好
	mergeSort(data, mid+1, hi)
	merge(data, lo, mid, hi)
}

func merge(data []int, lo, mid, hi int)  {
	subData1 := data[lo:mid+1]
	subData2 := data[mid+1:hi+1]
	// 额外申请辅助空间
	tmp := make([]int, hi-lo+1)
	p1 := 0
	p2 := 0
	index := 0 

	for p1<len(subData1) && p2<len(subData2) {
		if subData1[p1] < subData2[p2] {
			// tmp = append(tmp, subData1[p1])
			tmp[index] = subData1[p1]
			p1++
		} else {
			tmp[index] = subData2[p2]
			p2++
		}
		index++
	}



	// 将p1或者p2的剩余部分添加到tmp中
	for p1<len(subData1) {
		tmp[index] = subData1[p1]
		p1++
		index++
	}
	for p2<len(subData2) {
		tmp[index] = subData2[p2]
		p2++
		index++
	}


	// 将tmp中的数据还原到原数组data中
	for k:=0;k<len(tmp);k++ {
		data[lo+k] = tmp[k]		
	}
}


func main()  {
	s := []int{12,34,43,85,57,90,46}
	mergeSort(s, 0, len(s)-1)
	fmt.Println(s)
}