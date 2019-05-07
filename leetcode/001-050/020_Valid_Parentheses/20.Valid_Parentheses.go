package main

import "fmt"

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
