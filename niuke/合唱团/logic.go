package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, d int
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		nl := make([]int, 0, n)
		for n > 0 {
			n--
			var i int
			fmt.Scan(&i)
			nl = append(nl, i)
		}
		sort.Ints(nl)
		fmt.Scan(&k, &d)
		r := nl[len(nl)-1]
		k1 := 1
		for i := len(nl) - 2; i >= 0 && k1 != k; i-- {
			if nl[i+1]-nl[i] > d {
				r = nl[i]
				k1 = 1
				continue
			}
			r *= nl[i]
			k1++
		}
		fmt.Println(r)
	}
}
