package exercise

// 使用hash_table

func intersection(nums1 []int, nums2 []int) []int {
    var res []int
    resMap := make(map[int]struct{})
    m := make(map[int]struct{})
    // 对num1做去重，结果放在map里面（m）
    for _, value1 := range nums1 {
        if _, ok := m[value1]; !ok {
            m[value1] = struct{}{}
        }
    }
    // 把num1与num2相比较，重叠元素放在新的map中(resMap)
    for _, value2 := range nums2 {
        if _, ok := m[value2]; ok {
            resMap[value2] = struct{}{}
        }
    }
    // 把resMap中的元素转换成返回值类型
    for key, _ := range resMap {
        res = append(res, key)
    }
    return res
}