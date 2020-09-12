package main

import "fmt"

func main()  {
	n:=2
	for n>0{
		n--
		var s1   string
		fmt.Scan(&s1)
		l1:=len(s1)
		for i:=0;i<l1 ;i+=8{
			if i+8>l1 {
				fmt.Print(s1[i:l1])
				for j:=l1;j<i+8;j++{
					fmt.Print(0)
				}
				fmt.Println()
			}else {
				fmt.Println(s1[i:i+8])
			}
		}
	}

}
