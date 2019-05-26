package main

import "fmt"

func checkPalindrome(inputString string) bool {
	palindrome := ""
	for i := len(inputString) - 1; i >= 0; i-- {
		palindrome += string(inputString[i])
	}
	return palindrome == inputString
}

func checkPalindromeOptimize(s string) bool {
	return s == Reverse(s)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(checkPalindromeOptimize("aabaa") == true)
	fmt.Println(checkPalindrome("abac") == false)
	fmt.Println(checkPalindrome("a") == true)
	fmt.Println(checkPalindrome("az") == false)
	fmt.Println(checkPalindrome("abacaba") == true)
}
