package main

import "fmt"

func main()  {
		var s1   string
		fmt.Scan(&s1)

		var r int
	for i,v :=range s1{
		if i<2{
			continue
		}
		if v>='A' {

			r =r *16+ int(v-'A')+10
		}else {

			r =r *16+ int(v-'0')

		}
	}
	fmt.Println(r)
}
