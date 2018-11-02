// https://leetcode.com/problems/median-of-two-sorted-arrays/description/

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	r := combine(nums1, nums2)
	return media(r)
}

func combine(num1, num2 []int) []int {
	len1, i := len(num1), 0
	len2, j := len(num2), 0
	r := make([]int, len1+len2)
	for k := 0; k < len(r); k++ {
		if i < len1 && j < len2 {
			if num1[i] < num2[j] {
				r[k] = num1[i]
				i++
			} else {
				r[k] = num2[j]
				j++
			} 
			continue
		}
		if i < len1 {
			r[k] = num1[i]
			i++
		}
		if j < len2 {
			r[k] = num2[j]
			j++
		}
	}
	return r
}

func media(nums []int) float64 {
	l := len(nums)
	if l%2 == 0 {
		return float64(nums[l/2]+nums[l/2-1]) / 2.0
	}
	return float64(nums[l/2])
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	r := findMedianSortedArrays(nums1, nums2)
	fmt.Println(r)
}
