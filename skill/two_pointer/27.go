package twopointer

// leetcode27 移除元素
// https://leetcode.cn/problems/remove-element/

func removeElement(nums []int, val int) int {
	var rep int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[rep] = nums[i]
			rep++
			continue
		}
	}
	return rep
}
