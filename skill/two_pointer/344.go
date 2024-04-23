package twopointer

// leetcode 344 反转字符串
// 对撞指针的应用

func reverseString(s []byte) {
	// define inverse pointer which point to head and tail
	var head int
	tail := len(s) - 1
	for loop := 0; loop < (len(s) / 2); loop++ {
		// exchange
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
}
