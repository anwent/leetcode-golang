package main

import (
	"fmt"
)

type defInt int

// 查找二进制中1出现的次数
func (n defInt) findBits1() int {
	r := [] int{} 
	for n != 0 {
		if n % 2 == 1 {
			r = append(r, int(n))
		}
		n = n / 2
	}
	return len(r)
}

func countBits(num int) []int {
	r := [] int{}
    for i := 0; i <= num; i++ {
		num := defInt(i)
		r = append(r, num.findBits1())
	}
	return r
}

func main() {

	fmt.Println(countBits(5))
}


