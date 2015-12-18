package main
import "fmt"

var c int

func one_step(a,b int){
    if a == 20 && b == 20{
        c++
    } else {
        if a < 20{
            one_step(a+1, b)
        }
        if b < 20{
             one_step(a, b+1)
         }
    }
}
func problem15_slow()  {
    c = 0
    one_step(0,0)
    fmt.Println(c)
}

func pascal_triangle(n int) [][]int{
    var triangle [][]int
    var row = []int
    row = m

    return triangle
}

func problem15(){

}

func main() {
    problem15()
}
