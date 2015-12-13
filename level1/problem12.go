package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func NumDivisors(n int) int {
	l := common.PrimeFactors(n)
	//fmt.Println(n,l)
	freq_divisors := make([]int, len(l))

	act := l[0]
	idx := 0
	freq_divisors[0] = 1

	for i := 1; i < len(l); i++ {
		if l[i] != act {
			act = l[i]
			idx++
			freq_divisors[idx] = 1
		} else {
			freq_divisors[idx]++
		}
	}
	//fmt.Println(freq_divisors)
	p := 1
	return p
}

func problem12() {
	t := 1
	for i := 1; ; i++ {
		if NumDivisors(t) > 500 {
			fmt.Println("Problem 14:", t)
			break
		}
		t += i

	}
}
