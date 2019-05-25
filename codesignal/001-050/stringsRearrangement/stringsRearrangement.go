package main

import (
	"fmt"
)

func stringsRearrangement(inputArray []string) bool {
	return permutations(inputArray, func(array []string) bool {
		//fmt.Println(inputArray)
		before := array[0]
		for i := 1; i < len(array); i++ {
			end, mistakeNum := array[i], 1
			if before == end {
				return true
			}
			for j := range before {
				if before[j] != end[j] {
					mistakeNum--
					if mistakeNum < 0 {
						return false
					}
				}
			}
			before = end
		}
		return true
	})
}

func permutations(a []string, f func([]string) bool) bool {
	return permutation(a, f, 0)
}

func permutation(a []string, f func([]string) bool, i int) bool {
	if i > len(a) {
		f(a)
	}
	permutation(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		permutation(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
	return false
}

func main() {
	fmt.Println(stringsRearrangement([]string{"aba", "bbb", "bab"}))
	fmt.Println(stringsRearrangement([]string{"ab", "bb", "aa"}))
}
