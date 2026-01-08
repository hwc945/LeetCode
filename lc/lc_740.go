package lc

import "sort"

func DeleteAndEarn(nums []int) int {
	// 特殊情况
	if len(nums) == 0 {
		return 0
	}

	maxNum := 0
	numCount := make(map[int]int)
	// 统计每个数字的总和
	for _, num := range nums {
		numCount[num] += num
	}
	// 按照数字大小排序存储每个数字的总和
	dp := make([]int, len(numCount)+1)
	keys := make([]int, 0, len(numCount))
	for k := range numCount {
		keys = append(keys, k)
	}
	// 排序
	sort.Ints(keys)
	// 动态规划计算最大收益
	for i, key := range keys {
		if i == 0 {
			dp[i+1] = numCount[key]
		} else {
			if key == keys[i-1]+1 {
				dp[i+1] = max(dp[i], dp[i-1]+numCount[key])
			} else {
				dp[i+1] = dp[i] + numCount[key]
			}
		}
		maxNum = max(maxNum, dp[i+1])
	}
	return maxNum
}
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. 寻找最大值以确定桶的大小
	maxVal := 0
	for _, num := range nums {
		if num > maxVal {
			maxVal = num
		}
	}

	// 2. 构造桶：buckets[i] 代表数值为 i 的元素总分
	// 这一步直接消除了排序的需求
	buckets := make([]int, maxVal+1)
	for _, num := range nums {
		buckets[num] += num
	}

	// 3. 动态规划 (空间优化版：打家劫舍逻辑)
	// 我们只需要保存“上一个”和“上上一个”的最大值
	prev2 := 0 // 相当于 dp[i-2]
	prev1 := 0 // 相当于 dp[i-1]

	for _, score := range buckets {
		// 当前最大收益 = max(不选当前桶, 选当前桶 + 上上个桶的最大收益)
		temp := prev1
		if prev2+score > prev1 {
			prev1 = prev2 + score
		}
		prev2 = temp
	}

	return prev1
}
