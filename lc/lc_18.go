package lc

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
				}
				if sum > target {
					right--
				}
				if sum < target {
					left++
				}
			}
		}
	}
	return result
}

// 优化
func fourSum1(nums []int, target int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for a := 0; a < n-3; a++ { // 枚举第一个数
		x := nums[a]
		if a > 0 && x == nums[a-1] { // 跳过重复数字
			continue
		}
		if x+nums[a+1]+nums[a+2]+nums[a+3] > target { // 优化一
			break
		}
		if x+nums[n-3]+nums[n-2]+nums[n-1] < target { // 优化二
			continue
		}
		for b := a + 1; b < n-2; b++ { // 枚举第二个数
			y := nums[b]
			if b > a+1 && y == nums[b-1] { // 跳过重复数字
				continue
			}
			if x+y+nums[b+1]+nums[b+2] > target { // 优化一
				break
			}
			if x+y+nums[n-2]+nums[n-1] < target { // 优化二
				continue
			}
			c, d := b+1, n-1
			for c < d { // 双指针枚举第三个数和第四个数
				s := x + y + nums[c] + nums[d] // 四数之和
				if s > target {
					d--
				} else if s < target {
					c++
				} else { // s == target
					ans = append(ans, []int{x, y, nums[c], nums[d]})
					for c++; c < d && nums[c] == nums[c-1]; c++ {
					} // 跳过重复数字
					for d--; d > c && nums[d] == nums[d+1]; d-- {
					} // 跳过重复数字
				}
			}
		}
	}
	return
}
