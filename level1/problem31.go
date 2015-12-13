package main

import (
	"fmt"
)

var c int

func sum(coins, act [8]int) int {
	s := 0
	for i := 0; i < len(act); i++ {
		s += coins[i] * act[i]
	}
	return s
}

func backtrack(coins [8]int, act [8]int, index int, goal int) {
	// index es el indice de la moneda act
	// se incrementa act[index] o index

	if sum(coins, act) == goal {
		//fmt.Println(act)
		c += 1
		return
	}

	// possible decisions:
	// pick next coins
	if index > 0 {
		index--
		backtrack(coins, act, index, goal)
		index++
	}

	// add one more actual coin
	if sum(coins, act) < goal {
		act[index]++
		backtrack(coins, act, index, goal)
		act[index]--
	}
}

func problem31() {
	c = 0
	coins := [...]int{200, 100, 50, 20, 10, 5, 2, 1}
	act := [...]int{0, 0, 0, 0, 0, 0, 0, 0}
	backtrack(coins, act, 7, 200)
	fmt.Println(c)
}
