package main

import (
	"fmt"
	"math"
)

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{} //返回结果
	carryBit := 0         //进位标示
	for {
		var currentNode = result
		for {
			if currentNode.Next == nil {
				break
			}
			currentNode = currentNode.Next
		}
		currentSum := l1.Val + l2.Val + carryBit
		carryBit = 0 //进位归0
		if currentSum >= 10 {
			power := int(math.Pow(10, 1))
			currentNode.Val = (currentSum - currentSum/power*power) * 10 / power
			carryBit++
		} else {
			currentNode.Val = currentSum
		}
		if carryBit == 0 && l1.Next == nil && l2.Next == nil {
			break
		}
		currentNode.Next = &ListNode{Val: 0, Next: nil}
		if l1.Next == nil {
			l1 = &ListNode{Val: 0}
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2 = &ListNode{Val: 0}
		} else {
			l2 = l2.Next
		}
	}
	return result
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

	l1 = &ListNode{Next: l1}
	l2 = &ListNode{Next: l2}

	for *l1.Next != (ListNode{}) {
		l1 = l1.Next
		l2 = l2.Next

		sum := curr.Val + l1.Val + l2.Val
		mod := sum % 10

		curr.Val = mod
		curr.Next = &ListNode{}
		curr = curr.Next
		curr.Val = (sum - mod) / 10
	}

	return res
}

func addTwoNumbersOptimize1(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{} //返回结果
	current := &result
	carryBit := 0 //进位标示
	for l1 != nil || l2 != nil {
		var v1, v2 int
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		sum := v1 + v2 + carryBit
		current.Next = &ListNode{sum % 10, nil}
		carryBit = sum / 10
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carryBit != 0 {
		current.Next = &ListNode{1, nil}
	}
	return result.Next
}
func addTwoNumbersOptimize(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := ListNode{0, nil}
	carry := 0
	current := &dummyNode
	for l1 != nil || l2 != nil {
		var x, y int
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}
		sum := x + y + carry
		current.Next = &ListNode{sum % 10, nil}
		carry = sum / 10
		current = current.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry != 0 {
		current.Next = &ListNode{1, nil}
	}
	return dummyNode.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln ListNode) String() string {
	return fmt.Sprintf("{Val : %v,Next : %v}", ln.Val, ln.Next)
}

func main() {
	//num := 1526
	//n := 1
	//power := int(math.Pow(10, float64(n)))
	//result := (num - num/power*power) * 10 / power
	//fmt.Println(result)
	l1 := ListNode{Val: 5}
	l2 := ListNode{Val: 5}
	fmt.Println(addTwoNumbers(&l1, &l2))
	fmt.Println(addTwoNumbersOptimize(&l1, &l2))
	fmt.Println(addTwoNumbers1(&l1, &l2))
	l1 = ListNode{Val: 1, Next: &ListNode{Val: 8}}
	l2 = ListNode{Val: 0}
	fmt.Println(addTwoNumbers(&l1, &l2))
	fmt.Println(addTwoNumbersOptimize(&l1, &l2))
	fmt.Println(addTwoNumbers1(&l1, &l2))
	l1 = ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 = ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(&l1, &l2))
	fmt.Println(addTwoNumbersOptimize(&l1, &l2))
	fmt.Println(addTwoNumbers1(&l1, &l2))
}
