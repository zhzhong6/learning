package main
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main (){
	var n,m int
	var n1 string
	for {
		_,err:=  fmt.Scan(&n)
		if err!=nil{
			return
		}
		an :=make([]string,0)
		for n>0 {
			n--
			fmt.Scan(&n1)
			an =append(an, n1)
		}
		fmt.Scan(&m)
		am :=make(map[string]string,0)
		am2 :=make(map[string]int)
		ml :=make([]int,0)

		for m>0 {
			m--
			fmt.Scan(&n1)
			if _,ok :=am[n1];ok {
				continue
			}
			am [n1]=""
			nl1 ,_:=strconv.Atoi(n1)
			ml =append(ml,nl1 )

		}
		sort.Ints(ml)

		for k,v :=range an {
			for m1,_:=range am {
				if strings.Contains(v,m1){
					if len(am[m1])!=0{
						am[m1] +=" "
					}
					am[m1] += strconv.Itoa(k) + " "+ v
					am2[m1] +=1
				}
			}
		}
		var s string
		count :=0
		for _,m1:=range ml {
			ms := strconv.Itoa(m1)
			r :=am[ms]
			if len(r)==0 {
				continue
			}
			count += am2[ms]*2 + 2
			s+= ms+" "+strconv.Itoa(am2[ms])+" "+am[ms]
		}
		fmt.Println(count," ",s)
	}
	strings.Split()
}