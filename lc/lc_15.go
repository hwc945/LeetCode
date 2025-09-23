package lc

import "sort"

//给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请你返回所有和为 0 且不重复的三元组。
//注意：答案中不可以包含重复的三元组。

// 解题失败
func threeSum(nums []int) [][]int {
	result := [][]int{}
	i, j := 0, 0
	for i = 0; i < len(nums)-2; i++ {
		for j = i + 1; j < len(nums)-1; j++ {

		}
	}
	return result
}
func threeSum1(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	// 找出a + b + c = 0
	// a = nums[i], b = nums[j], c = -(a + b)
	for i := 0; i < len(nums); i++ {
		// 排序之后如果第一个元素已经大于零，那么不可能凑成三元组
		if nums[i] > 0 {
			break
		}
		// 三元组元素a去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		set := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			// 三元组元素b去重
			if j > i+2 && nums[j] == nums[j-1] && nums[j-1] == nums[j-2] {
				continue
			}
			c := -nums[i] - nums[j]
			if _, ok := set[c]; ok {
				res = append(res, []int{nums[i], nums[j], c})
				// 三元组元素c去重
				delete(set, c)
			} else {
				set[nums[j]] = struct{}{}
			}
		}
	}
	return res
}
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	// 找出a + b + c = 0
	// a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < len(nums)-2; i++ {
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重a
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
