package common

import "strconv"

func NumDigits(n int) int {
	s := strconv.Itoa(n)
	return len(s)
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, e := 0, len(s)-1; i < e; {
		if s[i] != s[e] {
			return false
		}
		i++
		e--
	}
	return true
}
