package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	reverse, origin := 0, x
	for x != 0 {
		val := x % 10
		reverse = reverse*10 + val
		x = x / 10
	}
	return origin == reverse
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
