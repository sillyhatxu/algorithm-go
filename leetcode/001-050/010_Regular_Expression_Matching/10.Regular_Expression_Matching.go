package main

import "fmt"

func isMatch(s string, p string) bool {
	if len(s) == 0 {
		return false
	} else if len(s) == 1 && len(p) == 1 {
		return s[0] == p[0] || p[0] == '.' || p[0] == '*'
	}
	//while (!s.empty() && (s[0] == p[0] || p[0] == '.')) {
	//	if (isMatch(s, p.substr(2))) return true;
	//	s = s.substr(1);
	//}
	//return isMatch(s, p.substr(2))
	return false
}

func main() {
	fmt.Println(isMatch("aa", "a") == false)
	fmt.Println(isMatch("aa", "a*") == true)
	fmt.Println(isMatch("ab", ".*") == true)
	fmt.Println(isMatch("aab", "c*a*b") == true)
	fmt.Println(isMatch("mississippi", "mis*is*p*.") == false)
}
