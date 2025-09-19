package lc

func twoSum(nums []int, target int) []int {
	aws := make([]int, 0)
	slow, fast := 0, 1
	for slow < len(nums) {
		if nums[slow]+nums[fast] == target {
			aws = append(aws, slow)
			aws = append(aws, fast)
			return aws
		}
		fast++
		if fast >= len(nums) {
			slow++
			fast = slow + 1
		}
	}
	return aws
}

type TwoSum struct {
	target   int
	position int
}

func twoSum1(nums []int, target int) []int {
	aws := make([]int, 0)
	num := make(map[int]TwoSum)
	for i, v := range nums {
		if val, ok := num[target-v]; ok && i != val.position {
			aws = append(aws, val.position)
			aws = append(aws, i)
			return aws
		}
		num[v] = TwoSum{target - v, i}
	}
	return aws
}

// 官方题解
func twoSum2(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
