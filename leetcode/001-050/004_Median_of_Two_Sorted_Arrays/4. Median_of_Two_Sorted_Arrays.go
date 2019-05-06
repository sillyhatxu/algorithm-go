package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size, i, j := len(nums1)+len(nums2), 0, 0
	slice := make([]int, size, size)
	for k := 0; k < size; k++ {
		if i > len(nums1)-1 && j <= len(nums2)-1 {
			slice[k] = nums2[j]
			j++
		} else if j > len(nums2)-1 && i <= len(nums1)-1 {
			slice[k] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			slice[k] = nums1[i]
			i++
		} else {
			slice[k] = nums2[j]
			j++
		}
	}
	if size%2 == 0 {
		return float64(slice[size/2-1]+slice[size/2]) / 2
	} else {
		return float64(slice[size/2])
	}
}

func main() {
	var array1 []int
	var array2 []int
	array1 = []int{}
	array2 = []int{1}
	fmt.Println(findMedianSortedArrays(array1, array2))
	array1 = []int{}
	array2 = []int{2, 3}
	fmt.Println(findMedianSortedArrays(array1, array2))
	array1 = []int{1, 2}
	array2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(array1, array2))
	array1 = []int{23, 29, 20, 32}
	array2 = []int{23, 21, 33, 25}
	fmt.Println(findMedianSortedArrays(array1, array2))
	array1 = []int{13, 22, 12}
	array2 = []int{8, 13}
	fmt.Println(findMedianSortedArrays(array1, array2))
}
