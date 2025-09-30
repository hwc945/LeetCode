package lc

import "strings"

// 利用双倍字符串特性判断是否由子串重复构成。
func repeatedSubstringPattern(s string) bool {
	if len(s) <= 1 {
		return false
	}
	doubled := s + s
	window := doubled[1 : len(doubled)-1]
	return strings.Contains(window, s)
}
