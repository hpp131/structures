package bubble

import "fmt"

// []int{12,34,43,85,57,90,46}

func bubbleSort(list []int) []int {
	n := len(list)
	if n <= 1 {
		return list
	}
	// 控制冒泡轮数
	for i:=n-1;i>0;i-- {
		// 控制每轮中的比较次数
		for j:=0;j<i;j++ {
			if list[j] > list[j+1] {
				// swith
				list[j],list[j+1] = list[j+1],list[j]
			}
		}
	}
	return list
}

func main()  {
	s := []int{12,34,43,85,57,90,46}
	fmt.Println(bubbleSort(s))
}