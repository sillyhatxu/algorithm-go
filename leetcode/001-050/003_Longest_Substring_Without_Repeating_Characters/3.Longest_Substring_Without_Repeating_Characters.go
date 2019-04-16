package main

import (
	"fmt"
)

/**
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

给定一个字符串，找到长度最长不重复的字符
翻译：这题看了好久没明白什么意思，直到看到了一个动态图，所以我自己重新做了一个图解，来帮助理解题意

*/

func lengthOfLongestSubstring(s string) int {
	maxLength := 0
	var currentArray []string
	for _, r := range s {
		c := string(r)
		for i, current := range currentArray {
			if current == c {
				currentArray = currentArray[i+1:]
			}
		}
		currentArray = append(currentArray, c)
		if maxLength < len(currentArray) {
			maxLength = len(currentArray)
		}
	}
	return maxLength
}

func Max(value1, value2 int) int {
	if value1 > value2 {
		return value1
	} else {
		return value2
	}
}

func lengthOfLongestSubstringOptimize(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := make(map[string]int)
	maxLength := 0
	j := 0
	for i, v := range s {
		c := string(v)
		if val, ok := m[c]; ok {
			j = Max(j, val+1)
		}
		m[c] = i
		maxLength = Max(maxLength, i-j+1)
	}
	return maxLength
}
func lengthOfLongestSubstringOptimizeBest(s string) int {
	var m [128]bool
	l := 0
	t := 0
	maxLength := 0
	for _, v := range s {
		c := v
		for m[c] {
			m[s[t]] = false
			t++
			l--
		}
		m[c] = true
		l++
		if maxLength < l {
			maxLength = l
		}
	}
	return maxLength
}

func main() {
	//test := []string{"abcabcbb", "bbbbb", "pwwkew", "ckilbkd"}
	test := []string{"pwwkew"}
	for _, src := range test {
		fmt.Println("lengthOfLongestSubstring : ", lengthOfLongestSubstring(src))
		fmt.Println("lengthOfLongestSubstringOptimize : ", lengthOfLongestSubstringOptimize(src))
		fmt.Println("lengthOfLongestSubstringOptimizeBest : ", lengthOfLongestSubstringOptimizeBest(src))
	}
}
