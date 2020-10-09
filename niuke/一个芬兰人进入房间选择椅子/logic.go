package main

import "fmt"

// 一维数组表示座位，1表示已经有人，0表示空着
func Choose(an []int) int {
	var  lastKey int
	amap := make(map[int]int)
	// 把连续空位个数记录下来  key：第一个空位的位置，val:空位的个数
	for i:=0;i<len(an);i++ {
		if an[i] == 1 {
			lastKey = i + 1
			continue
		}
		amap[lastKey]++
	}

	var max int
	// 找出最大的连续空位
	for k, v := range amap {
		if   amap[max] < v {
			max = k
		}
	}
	// 两端单独处理，两端只要靠一边坐就可以，
	if amap[0]>amap[lastKey] && amap[0]>(amap[max]+1)/2{
		return 0
	}
	if amap[lastKey]>amap[0] && amap[lastKey]>(amap[max]+1)/2{
		return len(an)-1
	}
	// 这里只返回一种情况，等价的情况只选一种，返回值是可以坐的位置
	return max+(amap[max])/2
}

func main() {
	an:=[]int{0,0,0,0,1,0,0,0,0,0,0,0,0,1,0}
	res:=Choose(an)
	fmt.Println(res)
}
