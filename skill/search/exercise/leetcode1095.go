package exercise


/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

 func findInMountainArray(target int, mountainArr *MountainArray) int {
    left, right := 0, mountainArr.length() - 1
    peak := SearchPeak(left, right, target, mountainArr)
    if first := binarySearch(left, peak); first != -1 {
        return first
    }else if second := binarySearch(peak + 1, mountainArr.length - 1); second != -1 {
        return second
    }else {
        return -1
    }
    
}   



// 1. 
func SearchPeak(left, right int, target int, mountainArr *MountainArray) int {
    for left <= right {
        mid := left + (right - left) / 2
        if mountainArr.get(mid) > mountainArr.get(mid + 1)  || mountainArr.get(mid) > mountainArr.get(mid - 1) {
            return mid
        // 在右半边
        }else if mountainArr.get(mid) > mountainArr.get(mid + 1){
            right = mid
        }else {
            left = mid
        }
    }
    return left
}
// 2. 两次二分
func binarySearch(start, end int, target int, mountainArr *MountainArray) int {
    left, right := start, end
    for left <= right {
        mid := left + (right - left) / 2
        if mountainArr.get(mid) == target {
            return mid
        }else if mountainArr.get(mid) > target {
            right = mid - 1
        }else {
            left = mid + 1
        }
    }
    return -1 
     
}