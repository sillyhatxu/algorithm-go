package main

import "fmt"

func mergeSort(array []int) []int {
	n := len(array)
	if n > 1 {
		mid := n / 2
		return merge(mergeSort(array[:mid]), mergeSort(array[mid:]))
	}
	return array
}

func merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	result := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			result[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			result[k] = left[i]
			i++

			//上边两个判断是某一个数组(left或right)已经把数值提取完毕，那么剩下的数组还有数值，所以只要依次取出剩下的那个数组数值即可
			//下边两个判断是比价数值
		} else if left[i] < right[j] {
			result[k] = left[i]
			i++
		} else { //left[i] > right[j]
			result[k] = right[j]
			j++
		}
	}
	return result
}

func main() {
	array := []int{}
	//fmt.Println(array)
	//fmt.Println(mergeSort(array))
	//fmt.Println("--------------")
	//array = []int{7, 5}
	//fmt.Println(array)
	//fmt.Println(mergeSort(array))
	//fmt.Println("--------------")
	//array = []int{82, 25, 59}
	//fmt.Println(array)
	//fmt.Println(mergeSort(array))
	//fmt.Println("--------------")
	//array = []int{82, 25, 59, 94}
	//fmt.Println(array)
	//fmt.Println(mergeSort(array))
	//fmt.Println("--------------")
	//array = []int{13, 14, 94, 33, 82, 25, 59, 94, 65, 23, 45, 27, 73, 25, 39, 10}
	//fmt.Println(array)
	//fmt.Println(mergeSort(array))
	//fmt.Println("--------------")
	array = []int{8, 7, 1, 9, 6, 2, 4, 5, 3}
	fmt.Println(array)
	fmt.Println(mergeSort(array))
}
