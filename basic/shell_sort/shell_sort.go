package main

import "fmt"

func shell_sort(array []int) []int {
	//for d := n / 2; d > 0; d /= 2 {}//原始
	//Hibbard  2^k - 1
	//Sedgewick
	//1, 8, 23, 77, 281,...
	//1, 5, 19, 41, 109,...
	d := 5
	n := len(array)
	for i := d; i < n; i += d {
		//j 表示当前比较的位置，如第一次i = 1 取出元素7，他需要与7这个元素之前的所有元素相比较，也就是循环 j
		for j := i; j >= d && j > 0; j -= d {
			if array[j] < array[j-d] { //交换位置
				array[j], array[j-d] = array[j-d], array[j]
			}
			fmt.Println(array)
		}
	}

	d = 1
	return array
}

func main() {
	array := []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	fmt.Println(array)
	fmt.Println(shell_sort(array))
}
