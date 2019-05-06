package main

import "fmt"

/**
20. Valid Parentheses
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

/**
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Valid Parentheses.
Next challenges:
*/
func isValid(s string) bool {
	signMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	index := -1
	var stack = make([]string, len(s))
	for _, v := range s {
		if index == -1 || signMap[string(v)] != stack[index] {
			index++
			stack[index] = string(v)
		} else {
			index--
		}
	}
	return index == -1
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("{[]}"))
}
