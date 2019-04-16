package main

import "fmt"

/**
4. Median of Two Sorted Arrays

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

有两个大小分别为m和n的排序数组nums1和nums2。
求两个排序数组的中位数(又叫中值 median)。总的运行时复杂度应该是O(log (m+n))。
您可以假设nums1和nums2不能同时为空。

中位数公式：
一组数组，按从小到大顺序排列  X[1] X[2] X[3] ... X[N]
则当N(数组的元素个数)为奇数时:
median = X[(N+1)/2]

当N为偶数时:
median = (X[N/2] + X[(N/2)+1])/2

例1
找出这组数据：23、29、20、32、23、21、33、25 的中位数。
解：
首先将该组数据进行排列（这里按从小到大的顺序），得到：
20、21、23、23、25、29、32、33
因为该组数据一共由8个数据组成，即n为偶数，故按中位数的计算方法，
得到中位数 median = (X[8/2] + X[(8/2)+1])/2
                = (X[4] + X[5/2])/2
                = (X[4] + X[5/2])/2   注意数学的 X1 是第一个元素，代码的X1是第二个元素
                = (23 + 25)/2
                = 24
即第四个数和第五个数的平均数。

例2
找出这组数据：10、20、 20、20、30的中位数。
解：
首先将该组数据进行排列（这里按从小到大的顺序），得到：
10、 20、 20、 20、 30
因为该组数据一共由5个数据组成，即n为奇数，故按中位数的计算方法，
得到中位数 median = X[(N+1)/2]
                = X[(5+1)/2]
                = X[3]
                = 20
即第3个数。

*/

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
