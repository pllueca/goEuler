package common

import "strconv"

func NumDigits(n int) int {
	s := strconv.Itoa(n)
	return len(s)
}

func SumDigits(n int) int {
	s := 0
	for n > 0 {
		s += n % 10
		n = n / 10
	}
	return s
}

func ReverseDigits(n int) int {
	r := 0
	for n > 0 {
		r *= 10
		r += n % 10
		n /= 10
	}

	return r
}

func IsPalindrome(n int) bool {
	s := strconv.Itoa(n)
	for i, e := 0, len(s)-1; i <= e; {
		if s[i] != s[e] {
			return false
		}
		i++
		e--
	}
	return true
}
