package main

import "fmt"

/**

14. Longest Common Prefix
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	result, check, index, circulation := "", "", 0, true
	for circulation {
		for i, v := range strs {
			if i == 0 {
				if len(v)-1 < index {
					circulation = false
					break
				} else {
					check = string(v[index])
				}
			} else {
				if len(v)-1 < index {
					circulation = false
					break
				} else if check != string(v[index]) {
					circulation = false
					break
				}
			}
		}
		if circulation {
			result += check
			index++
		}
	}
	return result
}

func main() {
	array := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(array) == "fl")
	array = []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(array) == "")
	array = []string{}
	fmt.Println(longestCommonPrefix(array) == "")
	array = []string{"aa", "a"}
	fmt.Println(longestCommonPrefix(array) == "a")
}
