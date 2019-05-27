package main

import "fmt"

func shapeArea(n int) int {
	if n == 1 {
		return n
	}
	return 4*(n-1) + shapeArea(n-1)
}

func shapeAreaOptimize(n int) int {
	return 2*n*n - 2*n + 1
}

func main() {
	fmt.Println(shapeAreaOptimize(1) == 1)
	fmt.Println(shapeAreaOptimize(2) == 5)
	fmt.Println(shapeAreaOptimize(3) == 13)
	fmt.Println(shapeAreaOptimize(4) == 25)
	fmt.Println(shapeAreaOptimize(5) == 41)
	fmt.Println(shapeAreaOptimize(7000) == 97986001)
}
