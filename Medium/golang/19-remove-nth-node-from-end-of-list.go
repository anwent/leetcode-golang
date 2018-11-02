// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/

// 参考 : https://blog.csdn.net/qq_20110551/article/details/81436531
package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 0  {
		return nil
	}
	// 创建一个头结点和快慢指针
	p_head := &ListNode{0, head}
	p_fast := p_head
	p_slow := p_head
	// 快指针先走n步
	for i := 0; i <= n; i++ {
		if p_fast == nil {
			return nil
		}
		p_fast = p_fast.Next
	}
	for p_fast != nil {
		p_fast = p_fast.Next
		p_slow = p_slow.Next
	}
	p_del := p_slow.Next
	p_slow.Next = p_del.Next
	return p_head.Next
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
}