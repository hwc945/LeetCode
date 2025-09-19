package lc

type XX struct {
	A int
	B int
}

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	k1, k2, k3, k4 := make(map[int]int), make(map[int]int), make(map[int]int), make(map[int]int)
	for _, v := range nums1 {
		k1[v]++
	}
	for _, v := range nums2 {
		for key, n := range k1 {
			k2[key+v] += n
		}
	}
	for _, v := range nums3 {
		for key, n := range k2 {
			k3[key+v] += n
		}
	}
	for _, v := range nums4 {
		for key, n := range k3 {
			k4[key+v] += n
		}
	}
	if v, ok := k4[0]; ok {
		count = v
	}
	return count
}

// 官方题解
func fourSumCount1(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}
