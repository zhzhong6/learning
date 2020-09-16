package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		if len(line) == 0 {
			break
		}
		ac := make(map[byte]int)
		for _, l := range line {
			ac[l] += 1
		}
		var min int
		for _, ac1 := range ac {
			if min == 0 || min > ac1 {
				min = ac1
			}
		}
		for _, ac1 := range ac {
			if min == 0 || min > ac1 {
				min = ac1
			}
		}
		r1 := make([]byte, 0)
		for _, l := range line {
			if ac[l] == min {
				continue
			}
			r1 = append(r1, l)
		}
		fmt.Println(string(r1))
	}
}
