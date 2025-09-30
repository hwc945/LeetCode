package lc

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//
// 每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	seen := make([]int, n+1)
	seen[1] = 1
	seen[2] = 2
	for i := 3; i <= n; i++ {
		seen[i] = seen[i-1] + seen[i-2]
	}
	return seen[n]
}
