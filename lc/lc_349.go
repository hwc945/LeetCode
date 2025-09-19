package lc

import "sort"

// 给定两个数组 nums1 和 nums2 ，返回 它们的 交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]int)
	var res []int
	for _, num := range nums1 {
		if _, ok := set[num]; ok {
		} else {
			set[num]++
		}
	}
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num)
		}
	}
	return res
}

// AI
func intersectionAI(nums1 []int, nums2 []int) []int {
	set := make(map[int]struct{})
	var res []int

	// 预分配结果切片的容量
	if len(nums1) < len(nums2) {
		res = make([]int, 0, len(nums1))
	} else {
		res = make([]int, 0, len(nums2))
	}

	// 存储nums1的元素
	for _, num := range nums1 {
		set[num] = struct{}{}
	}

	// 检查nums2的元素
	for _, num := range nums2 {
		if _, ok := set[num]; ok {
			res = append(res, num)
			delete(set, num) // 确保结果唯一
		}
	}

	return res
}

// 官方题解
// 两个集合
func intersection1(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}

// 排序 + 双指针
func intersection2(nums1 []int, nums2 []int) (res []int) {
	sort.Ints(nums1)
	sort.Ints(nums2)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		x, y := nums1[i], nums2[j]
		if x == y {
			if res == nil || x > res[len(res)-1] {
				res = append(res, x)
			}
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return
}
