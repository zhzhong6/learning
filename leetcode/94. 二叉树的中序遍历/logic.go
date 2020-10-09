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
	res:= inorderTraversal(root)
	fmt.Println(res)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var stackList *list.List
func inorderTraversal(root *TreeNode) []int {
	stackList = list.New()
	treeList := make([]int, 0)
	for stackList.Len() != 0 || root != nil {
		if root != nil {
			stackList.PushBack(root)
			root = root.Left
		} else {
			e := stackList.Back()
			root = e.Value.(*TreeNode)
			treeList = append(treeList, root.Val)
			stackList.Remove(e)
			root = root.Right
		}
	}
	return treeList
}