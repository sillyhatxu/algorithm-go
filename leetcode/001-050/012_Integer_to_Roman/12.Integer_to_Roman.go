package main

import (
	"fmt"
	"strings"
)

/**
Runtime: 12 ms, faster than 90.38% of Go online submissions for Integer to Roman.
Memory Usage: 3.4 MB, less than 36.54% of Go online submissions for Integer to Roman.
*/
func intToRoman(num int) string {
	switch {
	case num == 0:
		return ""
	case num >= 1000:
		return "M" + intToRoman(num-1000)
	case num >= 900:
		return "CM" + intToRoman(num-900)
	case num >= 500:
		return "D" + intToRoman(num-500)
	case num >= 400:
		return "CD" + intToRoman(num-400)
	case num >= 100:
		return "C" + intToRoman(num-100)
	case num >= 90:
		return "XC" + intToRoman(num-90)
	case num >= 50:
		return "L" + intToRoman(num-50)
	case num >= 40:
		return "XL" + intToRoman(num-40)
	case num >= 10:
		return "X" + intToRoman(num-10)
	case num >= 9:
		return "IX" + intToRoman(num-9)
	case num >= 5:
		return "V" + intToRoman(num-5)
	case num >= 4:
		return "IV" + intToRoman(num-4)
	default:
		return "I" + intToRoman(num-1)
	}
}

func intToRomanOptimize(num int) string {
	m, n, s, result := []string{"I", "V", "X", "L", "C", "D", "M", "", ""}, 1000, 3, ""
	for n > 0 {
		result += numToRoman(num/n, m[s*2], m[s*2+1], m[s*2+2])
		num -= num / n * n
		n = n / 10
		s--
	}
	return result
}

func numToRoman(num int, one string, five string, ten string) string {
	switch {
	case num == 9:
		return one + ten
	case num >= 5:
		return five + strings.Repeat(one, num-5)
	case num == 4:
		return one + five
	case num < 4:
		return strings.Repeat(one, num)
	default:
		return ""
	}
}

func main() {
	fmt.Println(intToRomanOptimize(400) == "CD")
	fmt.Println(intToRomanOptimize(4) == "IV")
	fmt.Println(intToRomanOptimize(9) == "IX")
	fmt.Println(intToRomanOptimize(40) == "XL")
	fmt.Println(intToRomanOptimize(58) == "LVIII")
	fmt.Println(intToRomanOptimize(1994) == "MCMXCIV")
}
