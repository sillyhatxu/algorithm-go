package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln ListNode) String() string {
	return fmt.Sprintf("{Val : %v,Next : %v}", ln.Val, ln.Next)
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	result := head.Next
	head.Next = swapPairs(head.Next.Next)
	result.Next = head
	return result
}

func main() {
	fmt.Println(swapPairs(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}))
	fmt.Println("2->1->4->3")
}
