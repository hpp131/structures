package main


/* 
题目描述

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。请你找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回 [-1, -1]。
你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。
*/


func searchRange(nums []int, target int) []int {
	start := binary(nums, target)
	if start >= len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := binary(nums, target+1) - 1
	return []int{start, end}
}



// 
func binary(data []int, target int) int {
	left, right := 0, len(data)-1
	for left <= right {
		mid := left + (right - left) / 2
		if target > data[mid] {
			left = mid + 1
		}else  {
			right = mid -1
		}
	}
	// 返回left可以保证left是第一个大于等于target的位置
	return left
}

/*
解题思路：
将题目转换为求第一个等于target的元素的下标位置（记作start），与第一个大于target的元素的下标位置（记作end），end-1即为最后一个target的位置。
但是二分搜索的写法需要稍做变更一下，不判断target = nums[mid]的情况，因为这样不能保证得出的结果就是第一个target的下标位置

*/
