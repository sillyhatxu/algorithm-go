package main

import (
	"fmt"
	"log"
)

func BubbleSortBasic(inputArray []int) ([]int, int, int) {
	timeComplexity := 0
	spaceComplexity := 0
	for range inputArray {
		for j := 0; j < len(inputArray)-1; j++ {
			timeComplexity++
			if inputArray[j] > inputArray[j+1] {
				temp := inputArray[j]
				inputArray[j] = inputArray[j+1]
				inputArray[j+1] = temp
			}
			fmt.Println(inputArray)
		}
	}
	return inputArray, timeComplexity, spaceComplexity
}

func BubbleSortBasicOptimize(inputArray []int) ([]int, int, int) {
	timeComplexity := 0
	spaceComplexity := 0
	for i := range inputArray {
		for j := 1; j < len(inputArray)-i; j++ {
			timeComplexity++
			if inputArray[j-1] > inputArray[j] {
				temp := inputArray[j-1]
				inputArray[j-1] = inputArray[j]
				inputArray[j] = temp
			}
			fmt.Println(inputArray)
		}
	}
	return inputArray, timeComplexity, spaceComplexity
}

func BubbleSortBasicOptimizeFlag(inputArray []int) ([]int, int, int) {
	timeComplexity := 0
	spaceComplexity := 0
	for i := range inputArray {
		flag := true //if no change;has been finished.
		for j := 1; j < len(inputArray)-i; j++ {
			timeComplexity++
			if inputArray[j-1] > inputArray[j] {
				flag = false
				temp := inputArray[j-1]
				inputArray[j-1] = inputArray[j]
				inputArray[j] = temp
			}
			fmt.Println(inputArray)
		}
		if flag {
			break
		}
	}
	return inputArray, timeComplexity, spaceComplexity
}

func main() {
	var output []int
	var timeComplexity int
	var spaceComplexity int

	output, timeComplexity, spaceComplexity = BubbleSortBasic([]int{9, 5, 1, 3, 4, 7, 6, 2, 8}) //Always O(N^2)  --> In fact, O(N * (N-1))
	log.Printf("Basic bubble sort : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)

	output, timeComplexity, spaceComplexity = BubbleSortBasicOptimize([]int{9, 5, 1, 3, 4, 7, 6, 2, 8})
	log.Printf("Basic optimize bubble sort : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)

	output, timeComplexity, spaceComplexity = BubbleSortBasicOptimizeFlag([]int{9, 5, 1, 3, 4, 7, 6, 2, 8})
	log.Printf("Basic optimize flag bubble sort : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)

	output, timeComplexity, spaceComplexity = BubbleSortBasicOptimizeFlag([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}) // O(N)
	log.Printf("The Best Situation : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)

	output, timeComplexity, spaceComplexity = BubbleSortBasicOptimizeFlag([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) // O(N^2)
	log.Printf("The Bad Situation : %v; Time Complexity : %v; Space Complexity : %v", output, timeComplexity, spaceComplexity)
}
