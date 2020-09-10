package main

import "fmt"

// 做不来
func main() {
	var n int
	fmt.Scan(&n)
	for n > 0 {
		n--
		var total, k, d1, d2 int
		fmt.Scan(&total, &k, &d1, &d2)
		if (total-(d1+d2+k))%3==0{
			fmt.Println("yes")
			return
		}

		d3:= d1-d2
		if d1<d2{
			d3=d2-d1
		}
		if (total-(d1+d3+k))%3==0{
			fmt.Println("yes")
			return
		}
		if (total-(d2+d3+k))%3==0{
			fmt.Println("yes")
			return
		}
		fmt.Println("no")
		return
	}

}
