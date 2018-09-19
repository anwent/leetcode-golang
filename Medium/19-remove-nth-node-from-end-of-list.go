// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/

package main

import (
	// "log"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p_fast := head
	p_slow := head
	index := 0
	for p_fast.Next != nil {
		p_fast = p_fast.Next
		if index >= n {
			p_slow = p_slow.Next
		}
		index++
	}
	p_slow.Next = p_fast
	return head
}

func main() {

	l1 := &ListNode{1, nil}
	l2 := &ListNode{2, nil}
	l3 := &ListNode{3, nil}
	l4 := &ListNode{4, nil}
	l5 := &ListNode{5, nil}

	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5



	r := removeNthFromEnd(l1, 2)

	fmt.Println(r)
	fmt.Println(r.Next)
	fmt.Println(r.Next.Next)
	fmt.Println(r.Next.Next.Next)




	age := 18

	ad_age := &age

	var p_ad_age *int = ad_age


	p_ad_age_1 := p_ad_age

	fmt.Println(age, ad_age, *p_ad_age, *p_ad_age_1)

}