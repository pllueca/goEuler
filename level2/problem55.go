package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func is_lychrel(n int) bool {
	for i := 0; i <= 50; i++ {
		n += common.ReverseDigits(n)
		if common.IsPalindrome(n) {
			return false
		}
	}
	return true
}
func main() {
	c := 0
	for i := 0; i < 10000; i++ {
		if is_lychrel(i) {
			c++
		}
	}
	fmt.Println("number of lychrel numbers below 10000: ", c)
}
