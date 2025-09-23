package lc

func reverseStr(s string, k int) string {

	tmp := []byte(s)
	n := len(s)
	for i := 0; i < n; i = i + 2*k {
		if i+k > n {
			reverseString(tmp[i:n])
		} else {
			reverseString(tmp[i : i+k])
		}
	}

	return string(tmp)
}
