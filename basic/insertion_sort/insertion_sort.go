package main

import "fmt"

func main() {
	array := []int{8, 7, 1, 9, 6, 2, 4, 5, 3}
	for i := 1; i < len(array); i++ {
		//j 表示当前比较的位置，如第一次i = 1 取出元素7，他需要与7这个元素之前的所有元素相比较，也就是循环 j
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] { //交换位置
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}
	fmt.Println(array)
}
