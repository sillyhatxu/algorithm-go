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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}
	p := head
	l := 0
	for p != nil {
		p = p.Next
		l++
	}
	n = l - n
	if n == 0 {
		return head.Next
	}
	q := head
	for q != nil {
		n--
		if n == 0 {
			if q.Next != nil {
				if q.Next.Next != nil {
					q.Next = q.Next.Next
				} else {
					q.Next = nil
				}
			} else {
				q.Next = nil
			}
		}
		q = q.Next
	}
	return head
}

func main() {
	fmt.Println(removeNthFromEnd(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}, 2))
}
