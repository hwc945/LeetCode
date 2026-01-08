package lc

func UniquePaths(m int, n int) int {
	// 定义一个长宽分别为m+1,n+1的二维切片

	ways := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		ways[i] = make([]int, n+1)
	}

	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if i == 1 && j == 1 {
				ways[i][j] = 1
			} else {
				ways[i][j] = ways[i-1][j] + ways[i][j-1]
			}
		}
	}
	return ways[m][n]
}
func uniquePaths(m int, n int) int {
	// 永远选择较小的作为列数，进一步压榨空间到 O(min(m, n))
	if m < n {
		m, n = n, m
	}

	dp := make([]int, n)
	// 初始化第一行全为 1（因为只能向右走，路径只有1条）
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	// 从第二行开始遍历
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 当前位置路径 = 上方路径 (dp[j]) + 左方路径 (dp[j-1])
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
