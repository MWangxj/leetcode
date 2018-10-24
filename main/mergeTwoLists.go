package main

import "fmt"

/**

Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

 */

//  Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var next *ListNode

	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val <= l2.Val {
				next = l1
				l1 = l1.Next
			} else {
				next = l2
				l2 = l2.Next
			}
		} else if l1 != nil {
			next = l1
			l1 = l1.Next
		} else {
			next = l2
			l2 = l2.Next
		}
		if tail != nil {
			tail.Next = next
			tail = next
		} else {
			head = next
			tail = next
		}
	}
	return head
}


func main() {
	res := mergeTwoLists(&ListNode{2, &ListNode{2, &ListNode{3, nil}}}, &ListNode{9, &ListNode{2, &ListNode{9, nil}}})
	fmt.Println(res)
}
