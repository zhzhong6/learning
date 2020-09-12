package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x   int
	fmt.Scan(&n, &x)
	an:=make([]int,n)
	min :=math.MaxInt32

	for i:=0;i<n;i++ {
		fmt.Scan(&an[i])
		if min>an[i]{
			min=an[i]
		}
	}
	count :=min*n
	for i:=0;i<n;i++ {
		 an[i]-=min
	}
	for {
		x= (x+n-1)%n
		if 	an[x]-1<0{
			an[x]=count
			break
		}
		count++
		an[x]--
	}
	for _,v:=range an {
		fmt.Print(v," ")
	}

}
