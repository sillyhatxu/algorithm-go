package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	previous := countAndSay(n - 1)
	result, check, length := "", string(previous[0]), 1
	for i := 1; i < len(previous); i++ {
		if check != string(previous[i]) {
			result += strconv.Itoa(length) + check
			check, length = string(previous[i]), 1
		} else {
			length++
		}
	}
	result += strconv.Itoa(length) + check
	return result
}

func countAndSayOptimize(n int) string {
	if n == 1 {
		return "1"
	}
	previous := countAndSay(n - 1)
	result, length := []byte{}, 1
	for i := 0; i < len(previous)-1; i++ {
		if previous[i] != previous[i+1] {
			result = append(result, []byte(strconv.Itoa(length))...)
			result = append(result, previous[i])
			length = 1
		} else {
			length++
		}
	}
	result = append(result, []byte(strconv.Itoa(length))...)
	result = append(result, previous[len(previous)-1])
	return string(result)
}

func main() {
	fmt.Println(countAndSayOptimize(1) == "1")
	fmt.Println(countAndSayOptimize(2) == "11")
	fmt.Println(countAndSayOptimize(3) == "21")
	fmt.Println(countAndSayOptimize(4) == "1211")
	fmt.Println(countAndSayOptimize(5) == "111221")
	fmt.Println(countAndSayOptimize(6) == "312211")
	fmt.Println(countAndSayOptimize(7) == "13112221")
	fmt.Println(countAndSayOptimize(8) == "1113213211")
	fmt.Println(countAndSayOptimize(9) == "31131211131221")
	fmt.Println(countAndSayOptimize(10) == "13211311123113112211")

	//fmt.Println(countAndSay(1) == "1")
	//fmt.Println(countAndSay(2) == "11")
	//fmt.Println(countAndSay(3) == "21")
	//fmt.Println(countAndSay(4) == "1211")
	//fmt.Println(countAndSay(5) == "111221")
	//fmt.Println(countAndSay(6) == "312211")
	//fmt.Println(countAndSay(7) == "13112221")
	//fmt.Println(countAndSay(8) == "1113213211")
	//fmt.Println(countAndSay(9) == "31131211131221")
	//fmt.Println(countAndSay(10) == "13211311123113112211")
}
