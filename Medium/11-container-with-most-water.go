// https://leetcode-cn.com/problems/container-with-most-water/description/

package main

import "fmt"

func maxArea(height []int) int {
	p_left := 0
	p_right := len(height)-1
	s_max := 0
	for p_left < p_right {
		s := 0
		if height[p_left] < height[p_right] {
			s = height[p_left] * (p_right - p_left)
			p_left++
		} else {
			s = height[p_right] * (p_right - p_left)
			p_right--
		}
		if s > s_max {
			s_max = s
		}
	}
	return s_max
}

func main() {

	h := []int{1,8,6,2,5,4,8,3,7,1}

	s := maxArea(h)

	fmt.Println(s)

}
