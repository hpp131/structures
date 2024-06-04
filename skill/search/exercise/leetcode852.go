package exercise

/*
题目描述：
符合下列属性的数组 arr 称为 山脉数组 ：
arr.length >= 3
存在 i（0 < i < arr.length - 1）使得：
arr[0] < arr[1] < ... arr[i-1] < arr[i]
arr[i] > arr[i+1] > ... > arr[arr.length - 1]
给你由整数组成的山脉数组 arr ，返回满足 arr[0] < arr[1] < ... arr[i - 1] < arr[i] > arr[i + 1] > ... > arr[arr.length - 1] 的下标 i 。
你必须设计并实现时间复杂度为 O(log(n)) 的解决方案。
*/

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	res := binarySearch(arr, left, right)
	return res
}

// 解法一：
func binarySearch(data []int, left, right int) int {
	for left <= right {
		mid := left + (right-left)/2
		// mid刚好是山峰，直接return
		if data[mid] > data[mid-1] && data[mid] > data[mid+1] {
			return mid
		}
		// mid位于山峰左边
		if data[mid] < data[mid+1] {
			left = mid
		}
		// mid位于山峰右边
		if data[mid] > data[mid+1] {
			right = mid
		}
	}
	return left
}

// 解法二：
// 该解法可以减少对原数组arr的引用。在特定的题目中对原数组arr的引用次数有次数限制，此时解法一无法通过。
func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	var ans int
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] > arr[mid+1] {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

/*
解题思路：找出最大值
*/
