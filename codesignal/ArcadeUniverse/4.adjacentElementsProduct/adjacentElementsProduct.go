package main

import "fmt"

func adjacentElementsProduct(inputArray []int) int {
	max := inputArray[0] * inputArray[1]
	for i := 1; i < len(inputArray)-1; i++ {
		temp := inputArray[i] * inputArray[i+1]
		if max < temp {
			max = temp
		}
	}
	return max
}

func main() {
	fmt.Println(adjacentElementsProduct([]int{3, 6, -2, -5, 7, 3}) == 21)
	fmt.Println(adjacentElementsProduct([]int{-1, -2}) == 2)
	fmt.Println(adjacentElementsProduct([]int{5, 1, 2, 3, 1, 4}) == 6)
	fmt.Println(adjacentElementsProduct([]int{1, 2, 3, 0}) == 6)
	fmt.Println(adjacentElementsProduct([]int{9, 5, 10, 2, 24, -1, -48}) == 50)
}
