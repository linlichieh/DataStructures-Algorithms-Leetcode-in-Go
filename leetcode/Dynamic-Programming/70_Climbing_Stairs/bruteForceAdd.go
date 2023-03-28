package leetcode70

// NOTE Bad solution
func bruteForce(n int) int {
	if n <= 3 {
		return n
	}
	var counter int
	bruteForceAdd(0, 0, n, &counter)
	return counter
}

func bruteForceAdd(num int, total int, target int, counter *int) {
	total += num
	if total == target {
		*counter++
		return
	} else if total > target {
		return
	}
	bruteForceAdd(1, total, target, counter)
	bruteForceAdd(2, total, target, counter)
}
