package exercise

import "fmt"

// 使用一个栈实现解决

func isValida(s string) bool {
    // 使用slice实现stack结构
    var sli []rune
    // 暂存前一个字符
    var tmp rune
    for _, value := range s {
        if value == '{' || value == '(' || value == '[' {
            sli = append(sli, value)
            tmp = value
        }else if len(sli)>0 && value == '}' && tmp == '{' {
            // pop stack
			// [:]的方式不会造成数组越界？sli[len(sli)-1]的方式会造成数组越界？
            sli = sli[:len(sli)-1]
			if len(sli)>0{
				tmp = sli[len(sli)- 1]
			}
        }else if len(sli)>0 && value == ')' && tmp == '(' {
            // pop stack
            sli = sli[:len(sli)-1]
			if len(sli)>0{
				tmp = sli[len(sli)- 1]
			}
        }else if len(sli)>0 && value == ']' && tmp == '[' {
            // pop stack
            sli = sli[:len(sli)-1]
			if len(sli)>0 {
				tmp = sli[len(sli)- 1]
			}        
		}else {
            return false
        }
    }
    fmt.Printf("sli len is %d, content is %[1]v", sli)
    if len(sli) != 0{
        return false
    }else {
        return true
    }
    
}