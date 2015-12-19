package level1

import (
	"fmt"
	"time"
)

func divisible(n int) bool {
	for i := 2; i <= 20; i++ {
		if n%i != 0 {
			return false
		}
	}
	return true
}

func Problem5() {
	n := 2520
	for {
		if divisible(n) {
			break
		}
		n += 2
	}
	fmt.Println("problem 5: ", n)
}

func main() {
	t0 := time.Now()
	Problem5()
	t1 := time.Now()
	fmt.Printf("Time: %v\n", t1.Sub(t0))
}
