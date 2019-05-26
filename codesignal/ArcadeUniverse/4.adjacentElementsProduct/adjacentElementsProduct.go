package main

import "fmt"

func adjacentElementsProduct(inputArray []int) int {

	return 0
}

func main() {
	fmt.Println(adjacentElementsProduct([]int{3, 6, -2, -5, 7, 3}) == 21)
	fmt.Println(adjacentElementsProduct([]int{-1, -2}) == 2)
	fmt.Println(adjacentElementsProduct([]int{5, 1, 2, 3, 1, 4}) == 6)
	fmt.Println(adjacentElementsProduct([]int{1, 2, 3, 0}) == 6)
	fmt.Println(adjacentElementsProduct([]int{9, 5, 10, 2, 24, -1, -48}) == 50)
}
