package exercise

import "fmt"

// leetcode 20
// 暴力解法

func isValid(s string) bool {
	loopNum := len(s)/2
    for n := 0;n < loopNum;n++ {
        for i:=0;i< len(s)-1;i++ {
            if s[i] == '(' && s[i+1] == ')' {
                s = s[:i] + s[i+2:]
                break
            }
            if s[i] == '[' && s[i+1] == ']' {
                s = s[:i] + s[i+2:]
                break
            }
            if s[i] == '{' && s[i+1] == '}' {
                s = s[:i] + s[i+2:]
                break
            }
        }
		fmt.Println("inner for loop exec")
    }
	fmt.Printf("s len is %d, s content is %s\n", len(s), s)
    if len(s) == 0 {
        return true
    }else {
        return false
    }
}

func main(){
	fmt.Println(isValid("()[]{}"))
}