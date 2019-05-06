package main

import "fmt"

func getValue(v byte) int {
	switch v {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

/**
Runtime: 16 ms, faster than 100.00% of Go online submissions for Roman to Integer.
Memory Usage: 3 MB, less than 22.00% of Go online submissions for Roman to Integer.
*/
func romanToInt(s string) int {
	h := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		current := h[s[i]]
		next := 0
		if i+1 < len(s) {
			next = h[s[i+1]]
		}
		if current >= next {
			result += current
		} else {
			result += next - current
			i++
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III") == 3)
	fmt.Println(romanToInt("IV") == 4)
	fmt.Println(romanToInt("IX") == 9)
	fmt.Println(romanToInt("LVIII") == 58)
	fmt.Println(romanToInt("MCMXCIV") == 1994)
}
