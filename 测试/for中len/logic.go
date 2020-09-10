package main

import (
	"fmt"
)

func main()  {
	l1:=[]int{1,2,3,4}

	// 这里len重新计算，每次循环结束重新算
	// 输出 0123456
	for i:=0;i<len(l1);i++ {
		if i<3 {
			l1=append(l1,5)
		}
		fmt.Print(i)
	}

	l2:=[]int{1,2,3,4}
	// range 这个就不会变
	// 输出 1234
	for i,v:=range l2{
		if i<3 {
			l2=append(l1,5)
		}
		fmt.Print(v)
	}
}

