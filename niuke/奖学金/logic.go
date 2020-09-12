package main

import (
	"fmt"
	"sort"
)

/*
5 10 9
0 5
9 1
8 1
0 1
9 100

*/

func main() {
	var n, r, avg int
	fmt.Scan(&n, &r, &avg)
	total := n * avg
	rl := make([]int, 0)
	rm := make(map[int]int)
	al := 0
	for n > 0 {
		n--
		var i, j int
		fmt.Scan(&i, &j)
		if _,ok:=rm[j];!ok{
			rl = append(rl, j)
		}
		rm[j] += r - i
		al += i
	}
	sort.Ints(rl)
	rt := 0
	for _, j := range rl {
		if total <= al {
			break
		}
		i := rm[j]
		if total-al <= i {
			rt += (total - al) * j
			break
		}
		rt += i * j
		al += i
	}
	fmt.Println(rt)
}
