package level1

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func Problem2() {
	fib := common.GenFib(1, 2)
	sum := 0
	var f int
	for i := 1; fib(i) < 4000000; i++ {
		f = fib(i)
		if f%2 == 0 {
			sum += f
		}
	}
	fmt.Println("problem 2:")
	fmt.Println(sum)
}
