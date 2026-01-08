package lc

func rotate(nums []int, k int) {
	k = k % len(nums)
	if k == 0 {
		return
	} else {
		left := nums[:len(nums)-k]
		right := nums[len(nums)-k:]
		copy(nums, append(right, left...))
	}

}

func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n

	reverse := func(s []int) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}

	reverse(nums)     // 第一步：全翻转
	reverse(nums[:k]) // 第二步：翻转前k个
	reverse(nums[k:]) // 第三步：翻转后面剩下的
}
