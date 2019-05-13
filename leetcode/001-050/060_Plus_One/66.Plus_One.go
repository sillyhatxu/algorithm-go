package main

import "fmt"

func plusOne(digits []int) []int {
	temp := digits[len(digits)-1] + 1
	if temp < 10 {
		digits[len(digits)-1] = temp
		return digits
	}
	carry := 1
	digits[len(digits)-1] = 0
	for i := len(digits) - 2; i >= 0; i-- {
		temp = digits[i] + carry
		if temp < 10 {
			digits[i] = temp
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append(digits[:0], append([]int{1}, digits[0:]...)...)
}

func main() {
	array := []int{}
	array = []int{9, 9, 9}
	fmt.Println(plusOne(array))
	array = []int{1, 2, 3}
	fmt.Println(plusOne(array))
	array = []int{4, 3, 2, 1}
	fmt.Println(plusOne(array))
}
