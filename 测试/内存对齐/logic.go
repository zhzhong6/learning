package main

import (
	"fmt"
	"unsafe"
)

type  str1 struct {
	a byte
	c byte
}
type  str2 struct {
	d float64
	c byte
	b int32
	a int64

}


func main() {
	s1:=unsafe.Sizeof(str1{})
	s2:=unsafe.Sizeof(str2{})
	fmt.Println(s1,s2)
}