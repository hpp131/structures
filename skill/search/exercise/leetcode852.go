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
    left, right := 0, len(arr) - 1
    res := binarySearch(arr, left, right)
    return res
}

func binarySearch(data []int, left, right int) int {
    for left <= right {
        mid := left + (right - left) / 2
        // mid刚好是山峰，直接return
        if data[mid] > data[mid - 1] && data[mid] > data[mid + 1]  {
            return mid
        }
        // mid位于山峰左边
        if data[mid] < data[mid + 1] {
            left = mid 
        }  
        // mid位于山峰右边
        if data[mid] > data[mid + 1] {
            right = mid 
        }
    }
    return left
}

/*
解题思路：找出最大值
*/