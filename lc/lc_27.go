package lc

//给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
//假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：

// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k

// 利用copy
func removeElement(nums []int, val int) int {
	var expectedNums []int
	for _, num := range nums {
		if num != val {
			expectedNums = append(expectedNums, num)
		}
	}
	copy(nums, expectedNums)
	return len(expectedNums)
}

// 首尾双指针
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

// AI 快慢指针 保持元素顺序
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
