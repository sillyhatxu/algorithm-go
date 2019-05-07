package main

import "fmt"

func removeElement(nums []int, val int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main() {
	array := []int{3, 2, 2, 3}
	fmt.Println(removeElement(array, 3))
	array = []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(array, 2))
}
