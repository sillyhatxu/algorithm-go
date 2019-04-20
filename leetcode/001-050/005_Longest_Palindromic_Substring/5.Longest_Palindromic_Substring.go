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
func longestPalindromeLow(s string) string {
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

//Dynamic Programming
//使用动态规划来解决
//执行结果
/**
Runtime: 64 ms, faster than 38.83% of Go online submissions for Longest Palindromic Substring.
Memory Usage: 6.5 MB, less than 18.26% of Go online submissions for Longest Palindromic Substring.
Next challenges:
*/
func longestPalindromeOptimize(s string) string {
	if len(s) == 0 || len(s) == 1 || (len(s) == 2 && s[0] == s[1]) {
		return s
	}
	dynamicProgrammingArray := make([][]bool, len(s))
	left, right, length := 0, 0, 0
	for i := 0; i < len(s); i++ {
		temp := make([]bool, len(s))
		dynamicProgrammingArray[i] = temp
		dynamicProgrammingArray[i][i] = true
		for j := 0; j <= i; j++ {
			dynamicProgrammingArray[j][i] = s[i] == s[j] && (i-j < 2 || dynamicProgrammingArray[j+1][i-1])
			if dynamicProgrammingArray[j][i] && length < i-j+1 {
				length = i - j + 1
				left = j
				right = i
			}
		}
	}
	return s[left : right+1]
}

func main() {
	fmt.Println(longestPalindromeOptimize(""))
	fmt.Println(longestPalindromeOptimize("a"))
	fmt.Println(longestPalindromeOptimize("bb"))
	fmt.Println(longestPalindromeOptimize("ccc"))
	fmt.Println(longestPalindromeOptimize("babad"))
	fmt.Println(longestPalindromeOptimize("forgeeksskeegfor"))
	fmt.Println(longestPalindromeOptimize("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}
