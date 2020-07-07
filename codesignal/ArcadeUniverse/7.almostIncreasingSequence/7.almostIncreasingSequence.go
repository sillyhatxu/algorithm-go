package main

import "fmt"

func almostIncreasingSequence(sequence []int) bool {
	index, defeat := 0, false
	for i := 1; i < len(sequence); i++ {
		if sequence[index] >= sequence[i] {
			if defeat {
				return false
			}
			defeat = true
			if index == 0 {
				index = i
			}
		} else {
			index = i
		}
	}
	return true
}

func main() {
	fmt.Println(almostIncreasingSequence([]int{1, 2, 5, 3, 5}) == true)
	fmt.Println(almostIncreasingSequence([]int{1, 2, 1, 2}) == false)
	fmt.Println(almostIncreasingSequence([]int{10, 1, 2, 3, 4, 5}) == true)
	fmt.Println(almostIncreasingSequence([]int{0, -2, 5, 6}) == true)
	fmt.Println(almostIncreasingSequence([]int{1, 3, 2, 1}) == false)
	fmt.Println(almostIncreasingSequence([]int{1, 3, 2}) == true)
	fmt.Println(almostIncreasingSequence([]int{3, 6, 5, 8, 10, 20, 15}) == false)
	fmt.Println(almostIncreasingSequence([]int{1, 1, 2, 3, 4, 4}) == false)
	fmt.Println(almostIncreasingSequence([]int{1, 4, 10, 4, 2}) == false)
}
