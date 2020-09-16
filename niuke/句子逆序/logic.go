package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		s=strings.ReplaceAll(s,"\n","")
		r:= strings.Split(s," ")
		for i:=len(r)-1;i>=0;i-- {
			fmt.Print(r[i])
			if i>0{
				fmt.Print(" ")
			}else{
				fmt.Println()
			}
		}


	}
}