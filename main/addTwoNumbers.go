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
	add, index, res := 0, 0, &ListNode{}
	for l1.Next != nil || l2.Next != nil {
		sum := l1.Val + l2.Val + add
		if sum > 9 {
			add = 1
			sum /= 9
		} else {
			add = 0
		}
		assignByIndex(index, sum, res)
		index++
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1.Next == nil {

	} else {
	}
	return res
}

func assignByIndex(index, val int, l *ListNode) {
	if index == 0 {
		l.Val = val
		return
	}
	if l.Next == nil {
		l.Next = &ListNode{}
	}
	index -= 1
	assignByIndex(index, val, l.Next)
}

func main() {
	res := addTwoNumbers(&ListNode{1, &ListNode{2, &ListNode{3, nil}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}})
	fmt.Println(res)
}
