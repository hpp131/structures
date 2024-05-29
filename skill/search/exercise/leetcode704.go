package exercise

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    // 避免(left + right) / 2 造成越界的情况
    for left <= right {
        mid := left + (right - left) / 2;
        if nums[mid] == target {
            return mid
        }else if target > nums[mid] {
            left = mid + 1
        }else  {
            right = mid - 1
        }
    }
    return -1 
}

/*
  使用正常的二分查找写法
*/