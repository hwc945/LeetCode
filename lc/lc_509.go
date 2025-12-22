package lc

func Fib(n int) int {
	fibs := make(map[int]int)
	fibs[0], fibs[1] = 0, 1
	return fibs2(n, fibs)
}

func fibs2(n int, fibs map[int]int) int {
	if _, ok := fibs[n]; ok {
		return fibs[n]
	} else {
		fibs[n] = fibs2(n-1, fibs) + fibs2(n-2, fibs)
		return fibs[n]
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	p, q, r := 0, 0, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = p + q
	}
	return r
}
