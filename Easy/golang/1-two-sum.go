// https://leetcode.com/problems/two-sum/description/

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))
	for i, _ := range nums {
		if j, ok := hash[target-nums[i]]; ok {
			return []int{j, i}
		}
		hash[nums[i]] = i
	}
	return nil
}

func main() {
	ns := []int{2, 4, 6, 8, 10}
	r := twoSum(ns, 12)
	fmt.Println(r)
}
