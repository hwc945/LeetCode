package lc

// @lc code=start
func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	lps := buildLPS(needle)
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = lps[j-1]
		}
		if haystack[i] == needle[j] {
			j++
			if j == len(needle) {
				return i - j + 1
			}
		}
	}
	return -1
}

func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	length := 0
	for i := 1; i < len(pattern); i++ {
		for length > 0 && pattern[i] != pattern[length] {
			length = lps[length-1]
		}
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
		}
	}
	return lps
}

// @lc code=end

// KMP 实现，暴露给其他文件复用。
func strStrKMP(haystack string, needle string) int {
	return strStr(haystack, needle)
}
