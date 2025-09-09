package lc

import "sort"

// 给定一个含有 n 个正整数的数组和一个正整数 target 。
// 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
func minSubArrayLen(target int, nums []int) int {
	count, sum := len(nums)+1, 0
	slow, fast := 0, 0
	for slow < len(nums) {
		if sum < target && fast < len(nums) {
			sum = sum + nums[fast]
			fast++
		} else if sum < target && fast == len(nums) && slow == 0 {
			return 0
		} else if sum < target && fast == len(nums) {
			return count
		} else if sum >= target {
			num := fast - slow
			if num < count {
				count = num
			}
			sum = sum - nums[slow]
			slow++
		}

	}

	return count
}

// AI解法
// 优化 滑动窗口
func minSubArrayLen1(target int, nums []int) int {
	minLen := len(nums) + 1 // 初始化为不可能的最大值
	sum := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		sum += nums[right] // 扩展窗口

		// 收缩窗口，直到 sum < target
		for sum >= target {
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
			}
			sum -= nums[left] // 收缩窗口
			left++
		}
	}

	if minLen == len(nums)+1 {
		return 0 // 未找到符合条件的子数组
	}
	return minLen
}

// 前缀和 + 二分查找
// 时间复杂度 O(n log n)，适用于大数据量但允许更高时间复杂度的场景。
func minSubArrayLen2(target int, nums []int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}

	minLen := n + 1
	for i := 1; i <= n; i++ {
		// 在 prefixSum[0..i-1] 中找最大的 j 使得 prefixSum[i] - prefixSum[j] >= target
		j := sort.Search(i, func(j int) bool {
			return prefixSum[i]-prefixSum[j] >= target
		})
		if j < i {
			minLen = min(minLen, i-j)
		}
	}

	if minLen == n+1 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
