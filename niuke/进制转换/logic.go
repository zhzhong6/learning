package main

import "fmt"

func main() {
	for {
		var s1 string
		_, err := fmt.Scan(&s1)
		if err != nil {
			break
		}
		var r int
		for i, v := range s1 {
			if i < 2 {
				continue
			}
			if v >= 'A' {

				r = r*16 + int(v-'A') + 10
			} else {

				r = r*16 + int(v-'0')

			}
		}
		fmt.Println(r)
	}
}
