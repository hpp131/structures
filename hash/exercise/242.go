package exercise

// 使用hash_table
import "fmt"

func isAnagram(s string, t string) bool {
    sli :=  [26]int{}
    fmt.Println(sli)
    for _, rune1 := range s {
        sli[rune1 - 'a'] += 1
    }
    for _, rune2 := range t {
        sli[rune2 - 'a'] -= 1
     }
    for _, value := range sli {
        if value != 0 {
            return false
        }
    }
    return true
}