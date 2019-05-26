package main

import "fmt"

func checkPalindrome(inputString string) bool {
	palindrome := ""
	for i := len(inputString) - 1; i >= 0; i-- {
		palindrome += string(inputString[i])
	}
	return palindrome == inputString
}

func main() {
	fmt.Println(checkPalindrome("aabaa") == true)
	fmt.Println(checkPalindrome("abac") == false)
	fmt.Println(checkPalindrome("a") == true)
	fmt.Println(checkPalindrome("az") == false)
	fmt.Println(checkPalindrome("abacaba") == true)
}
