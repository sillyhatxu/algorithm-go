package main

import "fmt"

func switchLetter(letter byte) []string {
	switch {
	case letter == '2':
		return []string{"a", "b", "c"}
	case letter == '3':
		return []string{"d", "e", "f"}
	case letter == '4':
		return []string{"g", "h", "i"}
	case letter == '5':
		return []string{"j", "k", "l"}
	case letter == '6':
		return []string{"m", "n", "o"}
	case letter == '7':
		return []string{"p", "q", "r", "s"}
	case letter == '8':
		return []string{"t", "u", "v"}
	case letter == '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func merge(a1, a2 []string) []string {
	var result []string
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			result = append(result, v1+v2)
		}
	}
	return result
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	} else if len(digits) == 1 {
		return switchLetter(digits[0])
	}
	result := switchLetter(digits[0])
	for i := 1; i < len(digits); i++ {
		result = merge(result, switchLetter(digits[i]))
	}
	return result
}

func main() {
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("234"))
}
