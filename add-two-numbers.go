package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 新建一个链表
	h := &ListNode{0, nil}
	cur := h
	// 进位
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		cur.Next = &ListNode{sum % 10, nil}
		cur = cur.Next
	}
	return h.Next
}

func main() {

	list1_1 := &ListNode{4, nil}
	list1_2 := &ListNode{3, nil}
	list1_3 := &ListNode{6, nil}

	list2_1 := &ListNode{7, nil}
	list2_2 := &ListNode{3, nil}
	list2_3 := &ListNode{9, nil}


	list1_1.Next = list1_2
	list1_2.Next = list1_3

	list2_1.Next = list2_2
	list2_2.Next = list2_3

	r := addTwoNumbers(list1_1, list2_1)

	fmt.Println(r.Val)
	fmt.Println(r.Next.Val)
	fmt.Println(r.Next.Next.Val)
	fmt.Println(r.Next.Next.Next.Val)

}

