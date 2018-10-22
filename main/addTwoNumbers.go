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

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add, res := 0, &ListNode{}
	current := res
	for l1.Next == nil || l2.Next == nil {
		sum := l1.Val + l2.Val + add
		if sum > 9 {
			add = 1
			sum /= 9
		} else {
			add = 0
		}
		current.Val = sum
		current.Next=&ListNode{}
		current=current.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1.Next==nil {
		current.Next=l2
	}else {
		current.Next=l1
	}
	return res
}

func main()  {
	res := addTwoNumbers(&ListNode{1,&ListNode{2,&ListNode{3,nil}}},&ListNode{1,&ListNode{2,&ListNode{3,nil}}})
	fmt.Println(res)
}
