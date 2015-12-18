package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func max_sum_bottom_top(t [][]int) int {


	var level, max_sum_below int
	var max_sums [][]int
	depth := len(t)
	max_sums := append(max_sums, t[len(t)-1)])
	for level = len(t) - 2; t > 0; t++ {
		max_sums_act := make(int, len(t[level]))
		for i := 0; i < len(t[level]); i++ {
			max_sums_act[i] = max(max_s[][i], max_s[][i])
		}
	}
}

func problem18() {
	dat, _ := ioutil.ReadFile("input18.txt")
	text := string(dat)
	//fmt.Println(text)
	linies := strings.Split(text, "\n")
	//fmt.Println(len(linies))
	var t [][]int
	for _, linia := range linies {
		if len(linia) < 1 {
			continue
		}
		nums := strings.Split(linia, " ")
		s_nums := make([]int, 0)
		for _, num := range nums {
			num_val, _ := strconv.Atoi(num)
			s_nums = append(s_nums, num_val)
		}
		t = append(t, s_nums)
	}

	fmt.Println(max_sum_bottom_top(t))

}

func main() {
	problem18()
}
