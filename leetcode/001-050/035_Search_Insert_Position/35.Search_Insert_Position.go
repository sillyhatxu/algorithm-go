package main

import "fmt"

func searchInsert(nums []int, target int) int {
	previous := 0
	for i, v := range nums {
		if v >= target {
			return i
		}
		previous++
	}
	return previous
}

func main() {
	array := []int{}
	array = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(array, 5) == 2)
	array = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(array, 2) == 1)
	array = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(array, 7) == 4)
	array = []int{1, 3, 5, 6}
	fmt.Println(searchInsert(array, 0) == 0)
}
