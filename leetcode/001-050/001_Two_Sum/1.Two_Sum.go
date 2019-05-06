package main

import "fmt"

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
