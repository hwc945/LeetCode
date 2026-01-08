package lc

func majorityElement(nums []int) int {
	see := make(map[int]int)
	for _, num := range nums {
		see[num]++
		if see[num] > len(nums)/2 {
			return num
		}
	}
	return -1
}
