package leetcode509

func nonRecurFib(n int) int {
	s := []int{0, 1}
	for i := 2; i <= n; i++ {
		s = append(s, s[i-1]+s[i-2])
	}
	return s[n]
}
