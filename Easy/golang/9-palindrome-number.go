// https://leetcode-cn.com/problems/palindrome-number/description/

package main

import (
	"fmt"
	"strconv"
)


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	str := strconv.Itoa(x)
	for i := 0; i < len(str) - 1; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(isPalindrome(5678))
}