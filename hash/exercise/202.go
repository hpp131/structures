package exercise

// 使用hash_table

func isHappy(n int) bool {
    m := make(map[int]bool)
    for n != 1 && !m[n] {
        n, m[n] = sum(n), true
    }
    return n == 1
}

// 求某数的各位平方之和
func sum(n int) int {
    sum := 0
    for n > 0 {
        sum += (n % 10) * (n % 10)
        n = (n / 10)
    }
    return sum
}