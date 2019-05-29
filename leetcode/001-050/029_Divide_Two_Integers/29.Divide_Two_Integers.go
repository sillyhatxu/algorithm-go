package main

import "fmt"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	divisorNegative, dividendNegative := false, false
	if divisor < 0 {
		divisorNegative = true
		// convert to positive form
		divisor = int(4294967295-uint32(divisor)) + 1
	}
	if dividend < 0 {
		dividendNegative = true
		// convert to positive
		dividend = int(4294967295-uint32(dividend)) + 1
	}
	// Doing this gives us a speed boost. I found without it, the
	// speed is 8ms when we use dividend as the max
	max := 1
	for multiply(divisor, max) < dividend {
		max = multiply(max, 2)
	}
	// search for the correct answer between 0 and max
	ans := doBinarySearch(0, max, dividend, divisor)
	// convert the answer to negative if either the numerator or denominator was negative
	if (divisorNegative && !dividendNegative) || (dividendNegative && !divisorNegative) {
		ans = int(int32(4294967295-uint32(ans))) + 1
	}
	if ans >= 2147483648 {
		ans = 2147483648 - 1
	}
	return ans
}

func multiply(base, num1 int) int {
	if num1 == 0 {
		return 0
	}
	if num1 == 1 {
		return base
	}
	res := multiply(base, num1>>1)
	if num1&1 != 0 {
		return res + res + base
	} else {
		return res + res
	}
}

func doBinarySearch(low, max, dividend, divisor int) int {
	if low > max {
		return 0
	}
	if low+1 == max {
		ans := multiply(divisor, max)
		if ans > dividend {
			return low
		} else {
			return max
		}
	}
	mid := (low + max) >> 1
	ans := multiply(divisor, mid)
	if ans > dividend {
		return doBinarySearch(low, mid, dividend, divisor)
	} else if ans < dividend {
		return doBinarySearch(mid, max, dividend, divisor)
	}
	return mid
}

func main() {
	fmt.Println(divide(-2147483648, -1) == 2147483647)
	fmt.Println(divide(-111, -1) == 111)
	fmt.Println(divide(9, 3) == 3)
	fmt.Println(divide(1, 1) == 1)
	fmt.Println(divide(10, 3) == 3)
	fmt.Println(divide(7, -3) == -2)
}
