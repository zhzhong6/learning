package main

import (
	"container/list"
	"fmt"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: nil,
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 3,
				Left: nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	res:= preorderTraversal(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	l := list.New()
	l.PushBack(root)
	res := make([]int, 0)
	for l.Len() != 0 {
		e := l.Front()
		node := e.Value.(*TreeNode)
		l.Remove(e)
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		l.PushFront(node.Right)
		l.PushFront(node.Left)
	}
	return res
}
