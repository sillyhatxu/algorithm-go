package main

import "fmt"

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
