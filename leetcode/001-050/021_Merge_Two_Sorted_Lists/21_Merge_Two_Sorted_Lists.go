package main

import "fmt"

/**
21. Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln ListNode) String() string {
	return fmt.Sprintf("{Val : %v,Next : %v}", ln.Val, ln.Next)
}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{0, nil}
	currentNode := &result
	for l1 != nil || l2 != nil {
		if l1 == nil {
			currentNode.Next = &ListNode{l2.Val, nil}
			l2 = l2.Next
		} else if l2 == nil {
			currentNode.Next = &ListNode{l1.Val, nil}
			l1 = l1.Next
		} else {
			if l1.Val < l2.Val {
				currentNode.Next = &ListNode{l1.Val, nil}
				l1 = l1.Next
			} else {
				currentNode.Next = &ListNode{l2.Val, nil}
				l2 = l2.Next
			}
		}
		currentNode = currentNode.Next
	}
	return result.Next
}

func main() {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(mergeTwoLists(l1, l2))
	l1 = &ListNode{Val: 5, Next: nil}
	l2 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}}
	fmt.Println(mergeTwoLists(l1, l2))
}
