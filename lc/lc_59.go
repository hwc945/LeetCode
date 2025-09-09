package lc

// 给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
func generateMatrix(n int) [][]int {
	// 初始化 n x n 的矩阵
	answer := make([][]int, n)
	for i := range answer {
		answer[i] = make([]int, n)
	}
	turn := 1
	left, right := 0, 0
	width, height := n, n-1
	for i := 0; i < n*n; {
		if turn == 1 {
			for j := 0; j < width; {
				answer[left][right] = i + 1
				i++
				j++
				right++
			}
			turn = 2
			right--
			left++
			width--
		} else if turn == 2 {
			for j := 0; j < height; {
				answer[left][right] = i + 1
				i++
				j++
				left++
			}
			turn = 3
			left--
			right--
			height--
		} else if turn == 3 {
			for j := 0; j < width; {
				answer[left][right] = i + 1
				i++
				j++
				right--
			}
			turn = 4
			right++
			left--
			width--
		} else {
			for j := 0; j < height; {
				answer[left][right] = i + 1
				i++
				j++
				left--
			}
			turn = 1
			left++
			right++
			height--
		}

	}

	return answer
}

// AI解法
func generateMatrix2(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1

	for num <= n*n {
		// 从左到右填充上边
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// 从上到下填充右边
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// 从右到左填充下边
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// 从下到上填充左边
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
