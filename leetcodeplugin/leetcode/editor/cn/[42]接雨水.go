//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。 
//
// 
//
// 示例 1： 
//
// 
//
// 
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 
// 
//
// 示例 2： 
//
// 
//输入：height = [4,2,0,3,2,5]
//输出：9
// 
//
// 
//
// 提示： 
//
// 
// n == height.length 
// 0 <= n <= 3 * 104 
// 0 <= height[i] <= 105 
// 
// Related Topics 栈 数组 双指针 
// 👍 1806 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
//package cn
//
//import "fmt"

func trap(height []int) int {
	_,a := change(height, 0)
	return a
}

func change(hs []int, res int) ([]int, int) {
	l := len(hs) - 1
	right := l
	var left int
	if right < 1 {
		return []int{}, res
	}
	if hs[0] < hs[l] {
		for hs[left] <= hs[0] {
			res += hs[0] - hs[left]
			left++
		}
	} else {
		for right>=0 &&  hs[right] <= hs[l]  {
			res += hs[l] - hs[right]
			right--
		}
	}
	return change(hs[left:right+1], res)

}


//leetcode submit region end(Prohibit modification and deletion)
