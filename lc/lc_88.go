package lc

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 创建辅助数组
	nums := make([]int, m+n)

	// 定义指针
	i, j := 0, 0

	// 循环 m+n 次，填满辅助数组
	for index := 0; index < m+n; index++ {
		// 情况 1: nums1 已经取完了，只能取 nums2
		if i == m {
			nums[index] = nums2[j]
			j++
			// 情况 2: nums2 已经取完了，只能取 nums1
		} else if j == n {
			nums[index] = nums1[i]
			i++
			// 情况 3: 两个都没取完，比较大小，nums1 小
		} else if nums1[i] <= nums2[j] {
			nums[index] = nums1[i]
			i++
			// 情况 4: 两个都没取完，nums2 小
		} else {
			nums[index] = nums2[j]
			j++
		}
	}

	// 【关键步骤】将辅助数组的结果复制回 nums1
	copy(nums1, nums)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	// p1 指向 nums1 有效数据的末尾
	p1 := m - 1
	// p2 指向 nums2 的末尾
	p2 := n - 1
	// tail 指向 nums1 总长度的末尾（填坑的位置）
	tail := m + n - 1

	// 只要 nums2 还有数据，就需要继续合并
	// (如果 p1 先用完了，剩下的 nums2 直接填进去)
	// (如果 p2 先用完了，nums1 剩下的本身就在原地，不用动)
	for p2 >= 0 {
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
		tail--
	}
}
