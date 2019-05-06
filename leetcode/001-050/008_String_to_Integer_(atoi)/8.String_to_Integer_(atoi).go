package main

import (
	"fmt"
	"math"
	"strconv"
)

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
