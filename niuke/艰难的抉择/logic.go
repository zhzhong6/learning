package main

import "fmt"

func main() {
	var n, i, r int
	fmt.Scan(&n)

	for n > 0 {
		n--
		fmt.Scan(&i)
		for i%2==0{
			r++
			i=i/2
		}
	}
  fmt.Println(r)
}

