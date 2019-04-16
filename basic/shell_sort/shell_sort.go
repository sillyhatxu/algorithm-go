package main

import "fmt"

func shell_sort(array []int) []int {
	//for d := n / 2; d > 0; d /= 2 {}//原始
	//Hibbard  2^k - 1
	//Sedgewick
	//1, 8, 23, 77, 281,...
	//1, 5, 19, 41, 109,...
	var result []int
	d := 5
	n := len(array)
	for i := d; i < n; i++ {
		//j 表示当前比较的位置，如第一次i = 1 取出元素7，他需要与7这个元素之前的所有元素相比较，也就是循环 j
		for j := i; j >= d && j > 0; j-- {
			if array[j] < array[j-1] { //交换位置
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}

	d = 1
	return result
}

func main() {
	array := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	fmt.Println(shell_sort(array))
}
