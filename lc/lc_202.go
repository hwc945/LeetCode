package lc

// 编写一个算法来判断一个数 n 是不是快乐数。
//
// 「快乐数」 定义为：
//
// 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
// 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
// 如果这个过程 结果为 1，那么这个数就是快乐数。
// 如果 n 是 快乐数 就返回 true ；不是，则返回 false 。
func isHappy(n int) bool {
	notOneNums := make(map[int]bool)
	for notOneNums[n] == false {
		is, newN := isOne(n)
		if is {
			return true
		} else {
			notOneNums[n] = true
			n = newN
		}
	}
	return false
}
func isOne(n int) (bool, int) {
	sum := 0
	for n != 0 {

		sum = sum + (n%10)*(n%10)
		n = n / 10
	}

	return sum == 1, sum
}

// Ai
func isHappyAI(n int) bool {
	seen := make(map[int]struct{})

	for {
		// 计算数字的平方和
		sum := 0
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n /= 10
		}

		// 检查是否为快乐数
		if sum == 1 {
			return true
		}

		// 检测循环
		if _, ok := seen[sum]; ok {
			return false
		}

		// 记录当前数字并继续
		seen[sum] = struct{}{}
		n = sum
	}
}

// 官方解法
// 快慢指针法
func isHappy1(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 && slow != fast {
		slow = step(slow)
		fast = step(step(fast))
	}
	return fast == 1
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n % 10) * (n % 10)
		n = n / 10
	}
	return sum
}

// 数学
func isHappy2(n int) bool {
	cycle := map[int]bool{4: true, 6: true, 37: true, 58: true, 89: true, 145: true, 42: true, 20: true}
	for n != 1 && !cycle[n] {
		n = step(n)
	}
	return n == 1
}
