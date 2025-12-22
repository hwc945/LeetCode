package lc

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	nums1 := make([]int, len(nums))
	nums1[0] = nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums1[count] = nums[i]
			count++
		}
	}
	copy(nums, nums1)
	return count
}
func RemoveDuplicates1(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	i, j := 0, 1
	for j < len(nums) {
		if nums[j] != nums[j-1] {
			nums[i+1] = nums[j]
			i++
		}
		j++
	}

	return i + 1
}
