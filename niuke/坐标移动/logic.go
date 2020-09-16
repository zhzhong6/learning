package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var x, y int
		sa := strings.Split(s, ";")
		for _, s1 := range sa {
			if len(s1) < 2 {
				continue
			}
			sn := s1[1:]
			r1 := changeNum(sn)
			switch s1[0] {
			case 'A':
				x -= r1
			case 'S':
				y -= r1
			case 'W':
				y += r1
			case 'D':
				x += r1
			}

		}
		output:=	strconv.Itoa(x)+","+strconv.Itoa(y)
		fmt.Println(output)
	}
}
func changeNum(s string) int {
	var r int
	for _, s1 := range s {
		if s1 > '9' || s1 < '0' {
			return 0
		}
		r = r*10 + int(s1-'0')
	}
	return r
}
