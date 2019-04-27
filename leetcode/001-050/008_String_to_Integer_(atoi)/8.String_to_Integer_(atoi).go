package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
8. String to Integer (atoi)
Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.
*/

/**
Runtime: 8 ms, faster than 15.82% of Go online submissions for String to Integer (atoi).
Memory Usage: 2.4 MB, less than 66.67% of Go online submissions for String to Integer (atoi).
Next challenges:
*/
func myAtoi(str string) int {
	resultSrc := ""
	sign := 1
	start := true
	for _, v := range str {
		if start {
			if v == ' ' {
				continue
			} else if v == '0' {
				start = false
				continue
			}
			_, err := strconv.Atoi(string(v))
			if v != '-' && v != '+' && err != nil {
				return 0
			}
			start = false
			if v == '-' {
				sign = -1
			} else if v == '+' {

			} else {
				resultSrc += string(v)
			}
		} else {
			_, err := strconv.Atoi(string(v))
			if err != nil {
				break
			}
			resultSrc += string(v)
		}
	}
	temp := ""
	check := true
	for _, v := range resultSrc {
		if check && v == '0' {
			continue
		}
		temp += string(v)
		check = false
	}
	resultSrc = temp
	if len(resultSrc) > 10 && sign > 0 {
		return math.MaxInt32
	} else if len(resultSrc) > 10 && sign < 0 {
		return -math.MaxInt32 - 1
	}
	r, err := strconv.ParseInt(resultSrc, 10, 64)
	if err != nil {
		return 0
	}
	result := int(r) * sign
	if result > math.MaxInt32 {
		return math.MaxInt32
	} else if result < -math.MaxInt32-1 {
		return -math.MaxInt32 - 1
	} else {
		return result
	}
}

/**
Runtime: 4 ms, faster than 100.00% of Go online submissions for String to Integer (atoi).
Memory Usage: 2.3 MB, less than 85.71% of Go online submissions for String to Integer (atoi).
*/
func myAtoiOptimize(str string) int {
	result, sign, start := 0, 1, false
	for _, v := range str {
		if v == ' ' && !start {
			continue
		} else if v == '+' && !start {
			sign = 1
			start = true
			continue
		} else if v == '-' && !start {
			sign = -1
			start = true
			continue
		}
		if '0' <= v && v <= '9' {
			start = true
			value, _ := strconv.Atoi(string(v))
			result = result*10 + value
			if result*sign > math.MaxInt32 {
				return math.MaxInt32
			} else if result*sign < -math.MaxInt32-1 {
				return -math.MaxInt32 - 1
			}
		} else {
			break
		}
	}
	return result * sign
}

func main() {
	array := []string{"-9223372036854775809", "0-1", "  0000000000012345678", "20000000000000000000", "+1", "   +0 123", "42", "    -42", "4193 with words", "words and 987", "-91283472332"}
	result := []int{-2147483648, 0, 12345678, 2147483647, 1, 0, 42, -42, 4193, 0, -2147483648}
	fmt.Println("---------- myAtoi ----------")
	for i, str := range array {
		fmt.Println(myAtoi(str) == result[i])
	}
	fmt.Println("---------- myAtoiOptimize ----------")
	for i, str := range array {
		fmt.Println(myAtoiOptimize(str) == result[i])
	}
}
