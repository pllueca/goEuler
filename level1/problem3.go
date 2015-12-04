package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func main() {

	n := 600851475143
	l := common.PrimeFactors(n)
	fmt.Println(l[len(l)-1])

}
