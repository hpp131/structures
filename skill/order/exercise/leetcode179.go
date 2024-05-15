package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

)

/*
给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。
注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。


示例 1：
输入：nums = [10,2]
输出："210"

示例 2：
输入：nums = [3,30,34,5,9]
输出："9534330"
*/

func main()  {
	exampleData := []int{1,2,3,4,5,6,7,8,9,0}
	strData := make([]string, 0)
	// 将元素类型由int转换为string
	for _, v := range exampleData {
		strData = append(strData, strconv.Itoa(v))
	}
	fmt.Println(strData)

	// 自定义比较各元素
	// 降序排序
	sort.Slice(strData, func(i, j int) bool {
		return strData[j] + strData[i] < strData[i] + strData[j]
	})
	// 此时strData中的元素已经是排序好的
	res := strings.Join(strData, "")
	if res[0] == '0' {
		return 
	}
	fmt.Println(res)

}

/* 题解：
本题虽然也属于排序算法，但是并没有手动的去实现某种排序算法，重要的是如何将源数组中的元素进行排序比较（这里需要根据题目要求进行自定义排序，sort.Slice即是
完成该作用）。
*/


