package main

import (
	"fmt"
	"github.com/brinon/goEuler/common"
	"math/big"
)

func main() {
	var a, b, max, maxa, maxb, limit, s int
	max = 0
	limit = 99

	for a = limit; a > 0; a-- {
		for b = limit; b > 0; b-- {
			biga, bigb := big.NewInt(a), big.NewInt(b)
			biga.Exp(biga, bigb , nil)
			s = common.BigSumDigits(biga)
			if s > max {
				max = s
				maxa, maxb = a,b
			}
		}
	}
	fmt.Println(max)
	fmt.Println(maxa, maxb)
}