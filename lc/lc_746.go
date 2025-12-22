package lc

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	for i := 0; i < len(cost); i++ {
		if i < 2 {
			dp[i] = cost[i]
		} else {
			if dp[i-1]+cost[i] < dp[i-2]+cost[i] {
				dp[i] = dp[i-1] + cost[i]
			} else {
				dp[i] = dp[i-2] + cost[i]
			}
		}
	}
	if dp[len(cost)-1] < dp[len(cost)-2] {
		return dp[len(cost)-1]
	} else {
		return dp[len(cost)-2]
	}
}
func minCostClimbingStairs1(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, m(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}

func m(a, b int) int {
	if a < b {
		return a
	}
	return b
}
