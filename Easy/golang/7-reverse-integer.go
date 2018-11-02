// https://leetcode.com/problems/reverse-integer/description/

package main

import (
	"math"
	"fmt"
)

func reverse(x int) int {
	var p int = 1
	if x < 0 {
		p = -1 
		x = x * -1
	}
	t := 0
	for x > 0 {
		t = t*10 + x%10
		x = x/10
	}
	if t > math.MaxInt32 || t < math.MinInt32 {
		t = 0
	}
	return t*p
}

func main() {
	// k := 123 % 10


	t := reverse(123)
	t1 := reverse(-123)
	t2 := reverse(120)


	fmt.Println(t, t1, t2)
}