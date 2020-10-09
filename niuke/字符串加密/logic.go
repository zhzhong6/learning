package main
import (
	"bufio"
	"fmt"
	"os"
)

func main (){
	reader := bufio.NewReader(os.Stdin)
	for {
		line ,_,_:=reader.ReadLine()
		if len(line)==0{
			return
		}
		fmt.Println(string(Encrypt(line)))
		line ,_,_ =reader.ReadLine()
		fmt.Println(string(unEncrypt(line)))

	}


}
func Encrypt ( s []byte)[]byte{
	r := make([]byte,0)
	for _,s1 :=range s {
		switch {
		case  s1>='0'&& s1<'9'  :
			r=append(r, s1+1)
		case s1=='9':
			r=append(r, '0')
		case s1=='z':
			r=append(r, 'A')
		case s1=='Z':
			r=append(r, 'a')
		case s1>='a' &&s1<'z':
			r=append(r, s1+1-'a'+'A')
		case s1>='A' &&s1<'Z':
			r=append(r, s1+1-'A'+'a')
		}
	}
	return r
}
func unEncrypt ( s []byte)[]byte{
	r := make([]byte,0)
	for _,s1 :=range s {
		switch {
		case  s1>'0'&& s1<='9'  :
			r=append(r, s1-1)
		case s1=='0':
			r=append(r, '9')
		case s1=='a':
			r=append(r, 'Z')
		case s1=='A':
			r=append(r, 'z')
		case s1>'a' &&s1<='z':
			r=append(r, s1-1-'a'+'A')
		case s1>'A' &&s1<='Z':
			r=append(r, s1-1-'A'+'a')
		}
	}
	return r
}
