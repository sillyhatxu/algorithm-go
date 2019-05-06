package main

import (
	"fmt"
	"math"
	"strconv"
)

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
