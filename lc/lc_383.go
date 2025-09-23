package lc

func canConstruct(ransomNote string, magazine string) bool {
	seen := make(map[rune]int)
	for _, v := range magazine {
		seen[v]++
	}
	for _, v := range ransomNote {
		if seen[v] > 0 {
			seen[v]--
		} else {
			return false
		}
	}
	return true
}
