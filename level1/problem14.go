package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func problem14() {
	max := 0
	maxi := -1
	var v int
	var lookup = map[int]int{1: 1}
	for i := 1000000; i > 0; i-- {
		v = common.DynamicCollatzSteps(i, lookup)
		if v >= max {
			max = v
			maxi = i
		}
	}
	fmt.Println(max, maxi)
}

func main() {

	problem14()
}
