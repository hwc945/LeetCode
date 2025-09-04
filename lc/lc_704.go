// Package lc leetcode算法题库解法，不借用AI
package lc

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果 target 存在返回下标，否则返回 -1。
// 你必须编写一个具有 O(log n) 时间复杂度的算法。
func search(nums []int, target int) int {
	position := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return position
}

// AI解法
func searchByAI(nums []int, target int) int {
	// 显式处理空数组的情况（可选，但有时更清晰）
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		// 防止整数溢出的标准写法
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid // 找到目标，直接返回索引
		} else if nums[mid] < target {
			left = mid + 1 // 目标在右半部分
		} else {
			right = mid - 1 // 目标在左半部分
		}
	}

	return -1 // 未找到目标
}

// 另一种写法，右开区间 [left, right)
func searchRight(nums []int, target int) int {
	left, right := 0, len(nums) // 注意：right 初始为 len(nums)，不是 len(nums)-1

	for left < right { // 因为右开，所以 left < right
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid // 因为右开，所以 right = mid 而不是 mid-1
		}
	}

	return -1
}
