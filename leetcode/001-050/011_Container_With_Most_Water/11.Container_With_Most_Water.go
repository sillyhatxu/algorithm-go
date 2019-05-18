package main

import "fmt"

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

/**
Runtime: 320 ms, faster than 24.67% of Go online submissions for Container With Most Water.
Memory Usage: 5.6 MB, less than 47.91% of Go online submissions for Container With Most Water.
*/
func maxArea(height []int) int {
	maxArea := 0
	for i, v := range height {
		for j := i + 1; j < len(height); j++ {
			currentArea := min(v, height[j]) * (j - i)
			if maxArea < currentArea {
				maxArea = currentArea
			}
		}
	}
	return maxArea
}

func maxAreaOptimize(height []int) int {
	maxArea, start, end := 0, 0, len(height)-1
	for start != end {
		currentArea := min(height[end], height[start]) * (end - start)
		if currentArea > maxArea {
			maxArea = currentArea
		}
		if height[end] > height[start] {
			start++ //The starting coordinate moves backwards
		} else {
			end-- //The end coordinate moves forward
		}
	}
	return maxArea
}

func main() {
	fmt.Println(maxAreaOptimize([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}) == 49)
}
