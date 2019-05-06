package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main() {
	array := []int{1, 1, 2}
	fmt.Println(removeDuplicates(array))
	array = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(array))
	array = []int{}
	fmt.Println(removeDuplicates(array))
}
