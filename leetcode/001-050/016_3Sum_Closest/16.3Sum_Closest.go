package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	near := math.MaxInt32
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			fmt.Println(nums[i], "     ", nums[left], "     ", nums[right])
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			} else if sum > target && sum-target < near {
				near = sum
				right--
			} else if sum < target && target-sum < near {
				near = sum
				left++
			} else {
				left++
			}
		}
	}
	return near
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1) == 2)
}
