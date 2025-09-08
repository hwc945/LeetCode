package lc

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

func sortedSquares(nums []int) []int {

	i, j, position := 0, len(nums)-1, len(nums)-1
	answer := make([]int, len(nums), len(nums))
	for i <= j {
		if i2, j2 := nums[i]*nums[i], nums[j]*nums[j]; i2 >= j2 {
			answer[position] = i2
			i++
		} else {
			answer[position] = j2
			j--
		}
		position--
	}
	return answer
}

// AI解法 双指针
func sortedSquares2(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	left, right := 0, n-1

	for k := n - 1; k >= 0; k-- {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare >= rightSquare {
			answer[k] = leftSquare
			left++
		} else {
			answer[k] = rightSquare
			right--
		}
	}
	return answer
}
