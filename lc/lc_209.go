package lc

import "sort"

// 滑动窗口在 O(n) 时间内找到满足条件的最短子数组。
func minSubArrayLen(target int, nums []int) int {
	minLen := len(nums) + 1
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for sum >= target {
			if current := right - left + 1; current < minLen {
				minLen = current
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0
	}
	return minLen
}

// 前缀和 + 二分查找：O(n log n)，适合需要演示另一种思路的场景。
func minSubArrayLenBinary(target int, nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	minLen := n + 1
	for i := 1; i <= n; i++ {
		j := sort.Search(i, func(j int) bool {
			return prefixSum[i]-prefixSum[j] >= target
		})
		if j < i && i-j < minLen {
			minLen = i - j
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}
