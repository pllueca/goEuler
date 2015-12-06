package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func problem4() {
	max := -1
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			if i*j > max && common.IsPalindrome(i*j) {
				max = i * j
			}
		}
	}
	fmt.Println(max)
}

func check(n int, c chan int) {
	if common.IsPalindrome(n) {
		c <- n
	} else {
		c <- -1
	}
}
func problem4chan() {
	c := make(chan int)
	max := -1
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			go check(i*j, c)
		}
	}
	var v int
	for v = <-c; ; v = <-c {
		if v > max {
			max = v
		}
	}

	fmt.Println(max)

}
