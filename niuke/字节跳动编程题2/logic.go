package main

import (
	"fmt"
)

func main() {
	var n, m, k ,r int
	var s string
	fmt.Scan(&n, &m, &s)
	l := make([]int, 0)
	for i := 1; i < n; i++ {
		if s[i] != s[i-1] {
			l = append(l, k+1)
			k = 0
		} else {
			k++
		}
	}

	for i := 1; i < len(l); i++ {
		t := m
		rl1 := 0
		for j := i; j < len(l); j += 2 {
			rl1 += l[j-1]
			if t-l[j] < 0 {
				rl1 += t
				break
			}
			rl1 += l[j]
			if t-l[j] == 0 && len(l)-1 > j {
				rl1 += l[j+1]
			}
			if t-l[j] == 0 {
				break
			}
			t -= l[j]
		}
		if r < rl1 {
			r = rl1
		}
	}

	fmt.Println(r)

}
