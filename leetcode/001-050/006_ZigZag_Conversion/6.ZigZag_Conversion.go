package main

import "fmt"

/**
6. ZigZag Conversion

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:

(you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

字符串"PAYPALISHIRING"根据给定行数，以一个Z字形显示出来，
如下所示：（您可能希望以固定字体显示此模式以获得更好的可读性）
其实例子太小了，大家看着不太方便看，所以我举一个超级大例子
Z             Z
Z           Z Z
Z         Z	  Z
Z       Z	  Z
Z     Z		  Z
Z   Z		  Z
Z Z			  Z
Z			  Z
这就是一个放倒了的 Z字母

制作多个图形得出公式  (数字间隔)M = 2N - 2

*/

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	if numRows == 2 {
		oddResult := ""
		evenResult := ""
		for i, v := range s {
			if i%2 == 0 {
				oddResult += string(v)
			} else {
				evenResult += string(v)
			}

		}
		return oddResult + evenResult
	}
	//var resultArray = make([]string, len(s))
	interval := 2*numRows - 2
	middleCount := numRows - 2
	middleNumber := make(map[int]bool)
	middleNumberIndex := numRows
	for middleNumberIndex < len(s) {
		nextIndex := middleNumberIndex + interval
		for i := 0; i < middleCount; i++ {
			middleNumber[middleNumberIndex] = true
			middleNumberIndex++
		}
		middleNumberIndex = nextIndex
	}
	result := ""
	for i := 0; i < numRows; i++ {
		index := i
		for index < len(s) {
			if i == 0 || i == numRows-1 {
				result += string(s[index])
			} else {
				result += string(s[index])
				intervalMax := index + interval
				if intervalMax > len(s) {
					intervalMax = len(s) - 1
				}
				for j := intervalMax; j > 0; j-- {
					if value, ok := middleNumber[j]; ok && value {
						result += string(s[j])
						middleNumber[j] = false
						break
					}
				}
			}
			index += interval
		}
	}
	return result
}

func main() {
	//fmt.Println(convert("0123456789", 3) == "0481357926")
	//fmt.Println(convert("PAYPALISHIRING", 0) == "")
	//fmt.Println(convert("PAYPALISHIRING", 1) == "PAYPALISHIRING")
	//fmt.Println(convert("ABCDEFG", 2) == "ACEGBDF")
	//fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("ABCDEFGHIJKLMNOPQ", 4))
	fmt.Println("AGMBFHLNCEIKOQDJP")
	fmt.Println(convert("ABCDEFGHIJKLMNOPQ", 4) == "AGMBFHLNCEIKOQDJP")
	fmt.Println(convert("PAYPALISHIRING", 4))
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
}
