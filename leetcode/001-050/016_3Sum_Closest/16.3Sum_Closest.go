package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	near, result := math.MaxInt32, math.MaxInt32
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			subTarget := target - sum
			if subTarget == 0 {
				return sum
			}
			if subTarget > 0 {
				left++
			} else {
				right--
				subTarget = -subTarget
			}
			if near < 0 {
				near = -near
			} else if near > subTarget {
				near = subTarget
				result = sum
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{82, -16, -60, -48, 32, -30, 0, -89, 70, -46, -82, -58, 66, 41, -96, -55, -49, -87, -33, -65, 9, 14, 81, -11, 80, -93, 79, -63, -61, -16, 22, -31, 75, 12, 17, -6, 37, -2, -89, -66, 3, -95, -74, 58, -95, 22, 11, -20, -36, 60, -75, 46, -52, -61, -28, 7, -50, -45, 93, -91, -43, 35, -99, -39, 53, -54, -98, -4, 13, -90, 23, -4, -65, 29, 85, -76, -64, 81, 32, -97, 51, 12, -82, -31, 81, -2, -83, -9, 89, -37, -23, -66, -91, -15, -98, -74, 94, 30, -2, -70, 13, 19, -77, -42, 33, -70, 25, 77, 47, -70, -70, 60, -63, -4, 83, 13, -78, -23, 28, -86, -11, -16, -57, -84, 51, -48, -63, -15, 29, 56, -25, 73, -69, 23, 28, 90, 96, 41, 65, -22, -43, -68, -77, 31, 69, -84, 23, -63, -18, 20, -93, -17, -48, -73, 14, -95, 98, -64, -12, -45, 14, 7, 51, -61, 89, -48, -23, 2, -56, 84, -2, 27, 74, -39, -18, -65, 79, -36, -76, -30, 44, 78, -76, 37, 88, 0, 32, 55, -51, 23, -9, 68, 26, 15, 66, 66, -56, -42, 56, 28, 33, -32, -23, -36, -12, -76, -12, 42, 12, 40, 69, 54, 82, -22, -7, 46, -46}, 270) == 270)
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1) == 2)
}
