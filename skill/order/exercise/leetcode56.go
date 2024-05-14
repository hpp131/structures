package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 对每个区间的start元素升序排列
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	tmpSlice := make([][]int, 0, len(intervals))
	tmpSlice = append(tmpSlice, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if tmpSlice[len(tmpSlice)-1][1] < intervals[i][0] {
			tmpSlice = append(tmpSlice, intervals[i])
		} else {
			tmpSlice[len(tmpSlice)-1][1] = max(tmpSlice[len(tmpSlice)-1][1], intervals[i][1])
		}
	}
	return tmpSlice

}

// 对每个区间的end元素进行升序;效果同merge函数相同
// 重点体会解法思想
func mergeReform(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	tmpSlice := make([][]int, 0, len(intervals))
	tmpSlice = append(tmpSlice, intervals[len(intervals)-1])
	for i := len(intervals) - 2; i >= 0; i-- {
		if tmpSlice[len(tmpSlice)-1][0] > intervals[i][1] {
			tmpSlice = append(tmpSlice, intervals[i])
		} else {
			tmpSlice[len(tmpSlice)-1][0] = min(tmpSlice[len(tmpSlice)-1][0], intervals[i][0])
		}
	}
	return tmpSlice

}

func main() {
	s := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	res := merge(s)
	resReform := mergeReform(s)
	fmt.Println(res)
	fmt.Println(resReform)
}

/*
在解题的时候并没有用到具体的某一种排序算法（如插入、快速、基数排序等，而是直接使用了Go语言内置sort包中的Slice方法，这个方法内部封装了各种排序算法。）
因此，这个题目不仅仅是一道考察排序的题目，更需要体会如下两种解题方式中所包含的思想：

按照start元素进行升序：
	使用sort.Slice进行排序后，每个区间的start元素都是有序（升序）的，此时第一个区间的start一定是整个interval数组中的最小元素，那我们只需要在
	if语句中比较每个区间的end元素即可

按照end元素进行升序：
	同理，使用sort.Slice进行排序后，每个区间的end元素都是有序（升序）的，此时我们区interval数组中的最后一个区间，并把这个区间放到tmpSlice数组
	中，作为tmpSlice的第一个元素，由于该元素中的end是整个interval中最大的元素，因此只需要在if语句中比较每个区间的start元素即可
*/
