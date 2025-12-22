package lc

import "sort"

func RemoveDuplicates0(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	keys := make(map[int]int, len(nums))
	count := len(nums)
	for i := 0; i < len(nums); i++ {
		keys[nums[i]]++
		if n, ok := keys[nums[i]]; ok && n >= 3 {
			nums[i] = nums[len(nums)-1] + 1
			count--
			continue
		}
	}
	sort.Ints(nums)
	return count
}
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	i, j := 2, 2
	for j < len(nums) {
		if nums[j] != nums[i-2] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
