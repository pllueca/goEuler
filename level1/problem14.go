package main

import (
	"github.com/brinon/goEuler/common"
	"fmt"
)

func problem14() {
	max := 0
	maxi := -1
	var v int
	var lookup = map[int]int{1:1}
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
	var lookup = map[int]int{1:0}
	fmt.Println(common.DynamicCollatzSteps(13, lookup))

	fmt.Println(common.DynamicCollatzSteps(524288, lookup))
	fmt.Println(common.DynamicCollatzSteps(837799, lookup))
	problem14()
}