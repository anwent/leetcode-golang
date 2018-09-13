// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/

/**
	需要优化  A
*/
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	head := 0
	footer := 0
	cur := make(map[string]int)
	max := 0
	for idx, _ := range s {
		head = idx
		cur[string(s[head])] = head
		for i := idx+1; i < len(s); i++ {
			footer = i
			if _, ok := cur[string(s[footer])]; ok {
				break
			}
			cur[string(s[footer])] = footer
		}
		if len(cur) > max {
			max = len(cur)
			fmt.Println("max len:", cur)
		}
		cur = map[string]int{}
	}
	return max
}

func main() {
	test := "pwwkew"
	length := lengthOfLongestSubstring(test)
	fmt.Println("max len is ", length)
}