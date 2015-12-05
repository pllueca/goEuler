package main

import (
	"fmt"

	"github.com/brinon/goEuler/common"
)

func problem3() {

	n := 600851475143
	l := common.PrimeFactors(n)
	fmt.Println("problem 3:")
	fmt.Println(l[len(l)-1])

}
