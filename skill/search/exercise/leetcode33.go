package main

import "fmt"

/*
题目描述

整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
你必须设计一个时间复杂度为 O(log n) 的算法解决此问题。

*/

func search(nums []int, target int) int {
	min := findK(nums)
	fmt.Println("k index is ", min)
	res1 := binarySearch(nums, 0, min-1, target)
	if res1 != -1 {
		return res1
	} else {
		res2 := binarySearch(nums, min, len(nums)-1, target)
		if res2 != -1 {
			return res2
		}
	}
	return -1
}

func findK(data []int) int {
	left, right := 0, len(data)-1
	for left <= right {
		mid := left + (right-left)/2
		// mid 在后半段
		if data[mid] == data[len(data)-1] {
			return len(data) - 1
		} else if data[mid] < data[len(data)-1] {
			right = mid - 1
			// mid 在前半段
		} else {
			left = mid + 1
		}
	}
	return left

}

// 最普通的二分查找操作
func binarySearch(data []int, left, right, target int) int {
	for left <= right {
		mid := left + (right-left)/2
		if data[mid] == target {
			return mid
		} else if data[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	s := []int{8, 9, 1, 2, 3, 4, 5, 6, 7}
	search(s, 5)
}

/*
解题思路:
先找出K所在的索引，k之前是一个有序数组，k之后也是一个有序数组，然后分别对两个数组使用二分查找。
因为题目规定了时间复杂度是O(logn),所以在找出K的步骤中，也应该使用二分来查找，直接遍历的话是O(n)!

*/

/*
总结:
做完几道二分查找的题目后，感觉二分查找的题目考察的是我们对于left、right指针、以及区间开闭的理解，这三个因素与最后return的值息息相关

*/
