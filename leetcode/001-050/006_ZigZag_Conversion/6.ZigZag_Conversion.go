package main

import "fmt"

func convertBackups(s string, numRows int) string {
	if len(s) == 0 || numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	if numRows == 2 {
		odd := ""
		event := ""
		for i, v := range s {
			if i%2 == 0 {
				odd += string(v)
			} else {
				event += string(v)
			}

		}
		return odd + event
	}
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

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	if numRows == 2 {
		odd := ""
		event := ""
		for i, v := range s {
			if i%2 == 0 {
				odd += string(v)
			} else {
				event += string(v)
			}
		}
		return odd + event
	}
	var result []byte
	rows, rIndex, upAndDown := make([][]byte, numRows), 0, false
	for i := range s {
		rows[rIndex] = append(rows[rIndex], s[i])
		if rIndex == numRows-1 || rIndex == 0 {
			upAndDown = !upAndDown
		}
		if upAndDown {
			rIndex++
		} else {
			rIndex--
		}
	}
	for i := range rows {
		result = append(result, rows[i]...)
	}
	return string(result)
}

func main() {
	fmt.Println(convert("0123456789", 3) == "0481357926")
	fmt.Println(convert("PAYPALISHIRING", 1) == "PAYPALISHIRING")
	fmt.Println(convert("ABCDEFG", 2) == "ACEGBDF")
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("ABCDEFGHIJKLMNOPQ", 4) == "AGMBFHLNCEIKOQDJP")
}
