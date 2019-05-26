package main

import "fmt"

func centuryFromYear(year int) int {
	if year%100 == 0 {
		return year / 100
	} else {
		return year/100 + 1
	}
}

func centuryFromYearOptimize(year int) int {
	return 1 + ((year - 1) / 100)
}

func main() {
	fmt.Println(centuryFromYear(1997) == 20)
	fmt.Println(centuryFromYear(1900) == 19)
}
