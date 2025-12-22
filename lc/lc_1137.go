package lc

func tribonacci(n int) int {
	if n < 2 {
		return n
	} else if n == 2 {
		return 1
	}
	i, j, k := 0, 1, 1
	for a := 3; a <= n; a++ {
		k, j, i = i+j+k, k, j
	}
	return k
}
