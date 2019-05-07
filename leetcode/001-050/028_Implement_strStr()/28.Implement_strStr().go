package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	length := len(needle)
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "bba"))
}
