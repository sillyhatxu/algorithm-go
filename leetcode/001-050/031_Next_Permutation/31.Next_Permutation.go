package main

import "fmt"

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					nums[i], nums[j] = nums[j], nums[i]
					return
				}
			}
		}
	}
}

func main() {
	var array []int
	array = []int{1, 3, 2}
	nextPermutation(array)
	fmt.Println(array)
	//array = []int{1, 2, 3}
	//nextPermutation(array)
	//fmt.Println(array)
}
