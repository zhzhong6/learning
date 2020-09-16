package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main (){
	reader:=bufio.NewReader(os.Stdin)
	for {

		s,err :=reader.ReadString('\n')
		if err!=nil{
			break
		}
		s =strings.ReplaceAll(s,"\n","")
		cm :=make(map[int]struct{})
		for _,s1 :=range s {
			switch   {
			case s1>='0'&& s1<='9':
				cm[1]=struct{}{}
			case s1>='a'&& s1<='z':
				cm[2]=struct{}{}
			case s1>='A'&& s1<='Z':
				cm[3]=struct{}{}
			default :
				cm[4]=struct{}{}
			}
		}

		tag :=true
		for i:=0;i<len(s)-2;i++ {
			var c int
			for j:=i+3;j<len(s)&& tag ;j++ {
				if s[i+c]==s[j] {
					c++
				}else{
					c=0
				}
				if c==3 {
					tag=false
					break
				}
			}
		}
		if tag && len(cm)>=3 && len(s)>8{
			fmt.Println("OK")
		}else {
			fmt.Println("NG")
		}

	}

}