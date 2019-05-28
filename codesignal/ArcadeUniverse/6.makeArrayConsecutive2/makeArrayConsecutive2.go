package main

import (
	"fmt"
)

func makeArrayConsecutive2(statues []int) int {
	if len(statues) == 1 {
		return 0
	}
	min, max := statues[0], statues[1]
	if min > max {
		min, max = max, min
	}
	for i := 2; i < len(statues); i++ {
		if statues[i] < min {
			min = statues[i]
		}
		if statues[i] > max {
			max = statues[i]
		}
	}
	return max - min - len(statues) + 1
}

func makeArrayConsecutive2Easy(statues []int) int {
	min, max := statues[0], statues[0]
	for i := 1; i < len(statues); i++ {
		if statues[i] < min {
			min = statues[i]
		}
		if statues[i] > max {
			max = statues[i]
		}
	}
	return (max - min + 1) - len(statues)
}

func main() {
	fmt.Println(makeArrayConsecutive2([]int{6, 2, 3, 8}) == 3)
	fmt.Println(makeArrayConsecutive2([]int{0, 3}) == 2)
	fmt.Println(makeArrayConsecutive2([]int{5, 4, 6}) == 0)
	fmt.Println(makeArrayConsecutive2([]int{6, 3}) == 2)
	fmt.Println(makeArrayConsecutive2([]int{1}) == 0)
}
