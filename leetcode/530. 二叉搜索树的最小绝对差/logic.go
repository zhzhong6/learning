package main

import (
	"fmt"
)

func main() {
	res := getMinimumDifference(&TreeNode{
		Val: 1,
		Left:nil,
		Right:&TreeNode{
			Val: 5,
			Left:&TreeNode{
				Val: 3,
				Left:nil,
				Right:nil,
			},
			Right:nil,
		},
	})
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	dif=0
	bianli(root)
	return dif
}

func bianli(root *TreeNode) {
	if root == nil {
		return
	}
	diff(root)
	if root.Left != nil {
		bianli(root.Left)
	}
	if root.Right != nil {
		bianli(root.Right)
	}
}

var dif int

func diff(n *TreeNode) {
	left := n.Left
	right := n.Right
	for left != nil {
		if (left.Right == nil && dif > n.Val-left.Val) || dif == 0 {
			dif = n.Val - left.Val
		}
		left = left.Right
	}
	for right != nil {
		if (right.Left == nil && dif > right.Val-n.Val) || dif == 0 {
			dif = right.Val-n.Val
		}
		right = right.Left
	}

}
