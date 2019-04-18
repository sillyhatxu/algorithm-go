package main

import "fmt"

/**
5. Longest Palindromic Substring

Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

给定一个字符串s，找出s中最长的回文子串。可以假设s的最大长度为1000。
题目简单，但是这里有一个定义

回文：正读反渎都相同的字符串，就是回文，如：aba -> aba;

*/

//最简单的理解方式，不过很现实时间复杂度太高了，LeetCode已经超时
func longestPalindromeLow1(s string) string {
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	var sonSrc []string
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			sonSrc = append(sonSrc, s[i:j])
		}
	}
	maxSrc := ""
	for _, v := range sonSrc {
		negate := ""
		for i := len(v) - 1; i >= 0; i-- {
			negate += string(v[i])
		}
		if v == negate && len(v) > len(maxSrc) {
			maxSrc = v
		}
	}
	return maxSrc
}

func longestPalindromeLow2(s string) string {
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	var sonSrc []string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := i; k <= k+j && k < len(s); k++ {
				fmt.Println(s[k])
			}
		}
		for j := len(s); j > i; j-- {
			sonSrc = append(sonSrc, s[i:j])
		}
	}
	maxSrc := ""
	for _, v := range sonSrc {
		negate := ""
		for i := len(v) - 1; i >= 0; i-- {
			negate += string(v[i])
		}
		if v == negate && len(v) > len(maxSrc) {
			maxSrc = v
		}
	}
	return maxSrc
}

func longestPalindrome(s string) string {
	return ""
}
func longestPalindromeOptimize(s string) string {
	if len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	var sonSrc []string
	for i := 0; i < len(s); i++ {
		for j := len(s); j > i; j-- {
			sonSrc = append(sonSrc, s[i:j])
		}
	}
	maxSrc := ""
	for _, v := range sonSrc {
		negate := ""
		for i := len(v) - 1; i >= 0; i-- {
			negate += string(v[i])
		}
		if v == negate && len(v) > len(maxSrc) {
			maxSrc = v
		}
	}
	return maxSrc
}

func test() {
	//maxSrc, positive, negative := "", "", ""
	//for i := 1; i < len(s); i++ {
	//	start := string(s[i])
	//	//tempMaxSrc := ""
	//	//check := false
	//	for j := i; j < len(s); j++ {
	//		end := string(s[j])
	//		positive += string(s[j])
	//		negative = string(s[j]) + negative
	//		if start == end {
	//			//check = true
	//			if positive != negative {
	//
	//			}
	//
	//			positive, negative = "", ""
	//			break
	//		}
	//		if len(positive) > len(maxSrc) {
	//			maxSrc = positive
	//		}
	//	}
	//
	//}
}

func main() {
	//fmt.Println(longestPalindrome("a"))
	//fmt.Println(longestPalindrome("bb"))
	//fmt.Println(longestPalindrome("ccc"))
	//fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindromeLow2("forgeeksskeegfor"))
	//fmt.Println(longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
