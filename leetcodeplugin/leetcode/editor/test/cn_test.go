package test

import (
	"fmt"
	"testing"
)

//func Test992(t *testing.T) {
//	sortArrayByParityII([]int{1,2,3,4})
//}
func Test42(t *testing.T) {
	res := reconstructQueue([][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
	fmt.Println(res)
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	res1 := make([]int, 0, k)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(res1); j++ {
			if nums1[i] > res1[j] && j+len(res1)-i < k {
				res1 = append(res1[:j], nums1[i])
				break
			}
			if len(res1) < k {
				res1 = append(res1, nums1[i])
			}
		}
	}
}
