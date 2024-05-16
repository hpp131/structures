package main

import (
	"fmt"

)

/*
给定一个包含红色、白色和蓝色、共 n 个元素的数组 nums ，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。

必须在不使用库内置的 sort 函数的情况下解决这个问题。
*/

// Solution1
func colorSort1(nums []int)  {
	var p int
	for i:=1;i<len(nums);i++ {
		if nums[i] == 0 {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	for j:=p;j<len(nums);j++ {
		if nums[j] == 1 {
			nums[j], nums[p] = nums[p], nums[j]
			p++
		}
	} 
}


// Solution2
func colorSort2(nums []int)  {
	var p0, p1 = 0, 0
	for i:=0;i<len(nums);i++ {
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[p1], nums[i] = nums[i], nums[p1]
			}
			p0++
			p1++
		} else if nums[i] == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}	
}






func main()  {
	nums := []int{1,0}
	// colorSort1(nums)
	colorSort2(nums)
	fmt.Println(nums)
}

/*
题解：
该题是排序题目，但是也没有使用到具体的某种排序算法（如快排 插入 选择 归并等），而是使用到了上述算法中的比较方法，即双指针/快慢指针的思想。
改题目写了两种题解，分别是colorSort1和colorSort2，两种方法中都蕴含快慢指针思想。
*/