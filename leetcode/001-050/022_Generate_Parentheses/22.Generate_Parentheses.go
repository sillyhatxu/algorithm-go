package main

import "fmt"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(result *[]string, str string, open, close, max int) {
	if len(str) == max*2 {
		*result = append(*result, str)
		return
	}
	if open < max {
		backtrack(result, str+"(", open+1, close, max)
	}
	if close < open {
		backtrack(result, str+")", open, close+1, max)
	}
}

func generateParenthesis1(n int) []string {
	res := make([]string, 0)
	var resT []byte
	var flags []int
	doGenereta(n, n, &res, resT, flags)
	return res
}

func doGenereta(n, availabe int, res *[]string, resT []byte, flags []int) {
	if len(resT) == 2*n {
		*res = append(*res, string(resT))
	} else {
		if availabe > 0 {
			doGenereta(n, availabe-1, res, append(resT, '('), append(flags, 0))
			if len(flags) > 0 {
				doGenereta(n, availabe, res, append(resT, ')'), flags[:len(flags)-1])
			}
		} else {
			doGenereta(n, availabe, res, append(resT, ')'), flags[:len(flags)-1])
		}
	}
}

func main() {
	fmt.Println(generateParenthesis(3))
	fmt.Println([]string{"((()))", "(()())", "(())()", "()(())", "()()()"})
}
