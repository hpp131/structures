package main

func maximumGap(nums []int) int {
    if len(nums) < 2 {
        return 0
    }
    

    maxNum := max(nums)
    for exp:=1;exp<=maxNum;exp*=10 {
        // 计数排序初始数组
        counter := make([]int, 10)
        for _, v := range nums {
            kNum := v / exp % 10
            counter[kNum]++
        }
        // 求前缀和
        for j:=1;j<len(counter);j++ {
            counter[j] += counter[j-1]
        }

        // 倒序遍历
        tmpSlice := make([]int, len(nums))
        for k:=len(nums)-1;k>=0;k-- {
            kNum := nums[k] / exp % 10
            tmpSlice[counter[kNum]-1] = nums[k]
            counter[kNum]--
        }
        copy(nums, tmpSlice)
    }

    var maxSep int
    for i:=1;i<len(nums);i++ {
        sep := nums[i] - nums[i - 1]
        if sep > maxSep {
            maxSep = sep
        }
        
    }
    return maxSep
}
    


func max(nums []int) int {
    var max int
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    return max

}


func main()  {
	s := []int{1,2,3,4,5,6,7,8,9,0}
	maximumGap(s)
}


/*
  该题要求使用线性复杂度和线性空间，而且明确指出需要先排序。那就是纯粹的排序问题，需要手写算法（基数排序）
*/