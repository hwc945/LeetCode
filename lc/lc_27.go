package lc

// 双指针原地移除目标元素，保持前缀顺序。
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

// 首尾双指针：不保留原始顺序，但交换次数更少。
func removeElement2(nums []int, val int) int {
	count := 0
	i, j := 0, len(nums)-1
	for i <= j {
		if val == nums[i] {
			count++
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}

	}
	return len(nums) - count
}

// 快慢指针：保留原始顺序。
func removeElement3(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
