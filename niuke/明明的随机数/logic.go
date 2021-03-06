package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, i int
	for {
		n1, err := fmt.Scan(&n)
		if err != nil || n1 == 0 {
			break
		}

		ma := make(map[int]struct{})
		l := make([]int, 0)
		for n > 0 {
			n--
			fmt.Scan(&i)
			if _, ok := ma[i]; ok {
				continue
			}
			ma[i] = struct{}{}
			l = append(l, i)
		}
		sort.Ints(l)
		for _, v := range l {
			fmt.Println(v)
		}
	}

}
