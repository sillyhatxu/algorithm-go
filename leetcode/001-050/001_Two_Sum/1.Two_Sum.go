package main

import "fmt"

/**

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

给定一个int数组，返回数组重点的两个数字加和等于目标target。
你可以假定每一次输入恰好有一个结果，你不可以使用相同的元素两次

*/
func twoSumBasic(nums []int, target int) []int {
	for i, n := range nums {
		for j, n2 := range nums {
			if i != j && n+n2 == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func twoSumOptimize(nums []int, target int) []int {
	temp := make(map[int]int)
	for i, n := range nums {
		_, hasValue := temp[n]
		if hasValue {
			return []int{temp[n], i}
		} else {
			temp[target-n] = i
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	fmt.Println(twoSumBasic(nums, 18))
	fmt.Println(twoSumOptimize(nums, 18))
}
