package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = x * -1
	}
	var result int64 = 0
	src := strconv.Itoa(x)
	for i := len(src) - 1; i >= 0; i-- {
		r := int64(x % 10)
		for j := 0; j < i; j++ {
			r = r * 10
		}
		result += r
		x = x / 10
	}
	if result > math.MaxInt32 || result < -math.MaxInt32-1 {
		return 0
	}
	return int(result) * sign
}

func main() {
	//fmt.Println(-123 % 10)
	array := []int{123, -123, 120, 1534236469}
	result := []int{321, -321, 21, 0}
	for i, str := range array {
		fmt.Println(reverse(str) == result[i])
	}
}
