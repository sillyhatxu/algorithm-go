package main

import (
	"fmt"
	"math"
	"strconv"
)

/**
7.Reverse Integer
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.
*/

func reverse(x int) int {
	if x > math.MaxInt32 || x < -math.MaxInt32-1 {
		return 0
	}
	check := true
	if x < 0 {
		x = x * -1
		check = false
	}
	src := strconv.Itoa(x)
	result := ""
	for i := len(src) - 1; i >= 0; i-- {
		result += string(src[i])
	}
	r, err := strconv.Atoi(result)
	if err != nil {
		return 0
	}
	if check {
		return int(r)
	} else {
		return int(r * -1)
	}
}

func main() {
	array := []int{123, -123, 120, 1534236469}
	result := []int{321, -321, 21, 9646324351}
	for i, str := range array {
		fmt.Println(reverse(str) == result[i])
	}
}
