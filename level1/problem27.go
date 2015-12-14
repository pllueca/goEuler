package level1

import (
	"fmt"
	"github.com/brinon/goEuler/common"
)

func main() {
	for i := 0; i <= 39; i++ {
		fmt.Println(i, common.FirstPrimes(i))
	}
}