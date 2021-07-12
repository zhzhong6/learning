package main

import (
	"fmt"
)

type str1 struct {
	a byte
	c byte
}
type str2 struct {
	d float64
	c byte
	b int32
	a int64
}

func isMatch(s string, p string) bool {
	fmt.Println(s, p)
	if (len(p) == 0 || p =="*" ) && len(s) == 0 {
		return true
	}

	if len(p) == 0 || (len(s) == 0 && p[0]!='*') {
		return false
	}

	switch p[0] {
	case '*':
		tag := 0
		stag := 0
		for i := 1; i < len(p); i++ {
			if p[i] == '*' {
				continue
			}
			if p[i] == '?' {
				stag++
				continue
			}
			tag = i
			break
		}

		res :=  isMatch("", p[tag:])
		for i := stag; i < len(s); i++ {
			if s[i] == p[tag] {
				res = res || isMatch(s[i:], p[tag:])
			}
		}
		return res
	case '?':
	default:
		if p[0] != s[0] {
			return false
		}
	}
	return isMatch(s[1:], p[1:])
}

func main() {

	r := isMatch("b", "*?*?")
	fmt.Println(r)
}
