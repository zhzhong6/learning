package main

import (
	"fmt"
)

func main() {
	res := canPartition([]int{100})
	fmt.Println(res)
}
func canPartition(nums []int) bool {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	db := make([][]bool, len(nums))
	for i := 0; i < len(db); i++ {
		db[i] = make([]bool, sum+1)
		db[i][0] = true
	}
	if nums[0]>sum {
		return false
	}
	db[0][nums[0]]=true
	for i := 1; i < len(db); i++ {
		for j := 1; j < sum+1; j++ {
			if  db[i-1][j] || (j>=nums[i] &&db[i-1][j-nums[i]] ) {
				db[i][j] =true
			}
		}
	}
	return 	db[len(db)-1][sum]
}
