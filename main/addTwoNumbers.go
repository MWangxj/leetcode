package main

import (
	`fmt`
)

/*

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.


*/


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add, res := 0, &ListNode{}
	current := res
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
		}
		if l2 != nil {
			b = l2.Val
		}
		sum := a + b + add
		if sum > 9 {
			add = 1
		} else {
			add = 0
		}
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if add > 0 {
		current.Next = &ListNode{Val: add}
	} else {
		current.Next = nil
	}
	return res.Next
}

func main() {
	res := addTwoNumbers(&ListNode{2, &ListNode{2, &ListNode{3, nil}}}, &ListNode{9, &ListNode{2, &ListNode{9, nil}}})
	fmt.Println(res)
}
