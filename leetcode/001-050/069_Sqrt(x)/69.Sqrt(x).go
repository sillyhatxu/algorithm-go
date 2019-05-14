package main

import "fmt"

func mySqrt(x int) int {
	current, previous := 0, x
	for true {
		square := current * current
		square1 := (current - 1) * (current - 1)
		square2 := (current + 1) * (current + 1)
		if x == square {
			return current
		} else if square > x && square1 < x {
			return current - 1
		} else if square < x && square2 > x {
			return current + 1
		}
		temp := current
		if square > x {

		} else {
			if current > previous {
				current = previous + (current-previous)/2
			} else {
				current = current + (previous-current)/2
			}
		}
		previous = temp
	}
	return 0
}

func main() {
	fmt.Println(mySqrt(1813156381))
	fmt.Println(mySqrt(4) == 2)
	fmt.Println(mySqrt(8) == 2)
}
