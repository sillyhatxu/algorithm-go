package main

import "fmt"

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	current, up, down := 0, x, 0
	for true {
		square := current * current
		square1 := (current - 1) * (current - 1)
		square2 := (current + 1) * (current + 1)
		if x == square || (square < x && square2 > x) {
			return current
		} else if square1 == x || (square > x && square1 < x) {
			return current - 1
		}
		if square > x {
			up = current
			current = down + (up-down)/2
		} else {
			down = current
			current = down + (up-down)/2
		}
	}
	return 0
}

func main() {
	fmt.Println(mySqrt(1) == 1)
	fmt.Println(mySqrt(1813156381) == 42581)
	fmt.Println(mySqrt(4) == 2)
	fmt.Println(mySqrt(8) == 2)
}
