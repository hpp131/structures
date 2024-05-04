package main

import (
	"fmt"

)

// example []int{12,34,43,85,57,90,46}

func sort(data []int, lo, hi int)  {
	if lo >= hi {
		return
	}
	// pivot := data[hi]
	p_index := partation(data, lo, hi)
	sort(data, lo, p_index-1)
	sort(data, p_index+1, hi)
}


// return pivot's index
func partation(data []int, lo, hi int) int {
	// pivot := data[hi]
	less := lo
	for great:=lo;great < hi;great++ {
		if data[great] < data[hi] {
			// switch less and great
			data[less], data[great] = data[great], data[less]
			less++
		}
	}
	// switch pivot and less index
	data[less], data[hi] = data[hi], data[less]
	return less
}

func main()  {
	quickS := []int{12,34,43,85,57,91,47}
	fmt.Printf("lenth is %d\n", len(quickS))
	sort(quickS, 0, len(quickS)-1)
	fmt.Println(quickS)
}