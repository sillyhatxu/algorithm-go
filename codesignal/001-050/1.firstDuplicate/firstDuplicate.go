package main

import "fmt"

func firstDuplicate(a []int) int {
	hash := make(map[int]bool)
	for _, v := range a {
		if _, ok := hash[v]; ok {
			return v
		} else {
			hash[v] = true
		}
	}
	return -1
}

func main() {
	fmt.Println(firstDuplicate([]int{2, 1, 3, 5, 3, 2}) == 3)
	fmt.Println(firstDuplicate([]int{8, 4, 6, 2, 6, 4, 7, 9, 5, 8}) == 6)
	fmt.Println(firstDuplicate([]int{2, 2}) == 2)
	fmt.Println(firstDuplicate([]int{2, 4, 3, 5, 1}) == -1)
	fmt.Println(firstDuplicate([]int{1}) == -1)
}
