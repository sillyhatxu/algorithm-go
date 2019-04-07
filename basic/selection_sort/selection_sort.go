package main

import (
	"fmt"
	"log"
)

func SelectionSort(inputArray []int) ([]int, int, int) {
	fmt.Println(inputArray)
	timeComplexity := 0
	spaceComplexity := 0
	for i, input := range inputArray {
		minIndex := i
		min := input
		for j := i; j < len(inputArray); j++ {
			timeComplexity++
			if min > inputArray[j] {
				min = inputArray[j]
				minIndex = j
			}
		}
		if min != input {
			inputArray[minIndex] = input
			inputArray[i] = min
		}
		fmt.Println(inputArray)
	}
	return inputArray, timeComplexity, spaceComplexity
}

func main() {
	var output []int
	var timeComplexity int
	var spaceComplexity int
	output, timeComplexity, spaceComplexity = SelectionSort([]int{9, 5, 1, 3, 4, 7, 6, 2, 8}) //Always O(N^2)  --> In fact, O(N * (N-1))
	log.Printf("Basic selection sort : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)

}
