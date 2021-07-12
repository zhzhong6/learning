package test

import (
	"fmt"
	"testing"
)

//func Test992(t *testing.T) {
//	sortArrayByParityII([]int{1,2,3,4})
//}
func Test42(t *testing.T) {
	res := p( []int{1, 2, 3})
	fmt.Println(res)
}


func permute (nums []int) [][]int {
	bn :=make([][]int,0)
	if len(nums)==1 {
		return [][]int{nums}
	}
	for i, n :=range nums {
		pn := append ([]int{},nums[:i]...)
		pn =append (pn,nums[i+1:]...)
		for _,p1 :=range  permute(pn) {
			pp :=append([]int{n},p1...)
			bn =append(bn,pp)
		}
	}
	return bn
}