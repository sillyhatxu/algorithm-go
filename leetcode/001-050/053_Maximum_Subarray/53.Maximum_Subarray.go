package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]
	for _, v := range nums[1:] {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if max < sum {
			max = sum
		}
	}
	return max
}

func main() {
	array := []int{}
	array = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(array) == 6)
	array = []int{1}
	fmt.Println(maxSubArray(array) == 1)
	array = []int{-2, 1}
	fmt.Println(maxSubArray(array) == 1)
}
