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

func postorderTraversal(root *TreeNode) []int {
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
		l.PushFront(node.Left)
		l.PushFront(node.Right)

	}
	res2 := make([]int,len(res))
	for i:=0;i<len(res);i++ {
		res2[i]=res[len(res)-1-i]
	}
	return res2
}


func postorderTraversal(root *TreeNode) []int {
	stackList := list.New()
	stackList.PushBack(root)
	treeList := make([]int, 0)
	if root==nil{
		return treeList
	}
	var pre * TreeNode
	for stackList.Len() != 0 {
		e:= stackList.Back()
		curr:=e.Value.(*TreeNode)
		if(curr.Left == nil && curr.Right == nil) || (pre != nil && (pre == curr.Left || pre == curr.Right)){
			//如果当前结点左右子节点为空  或  上一个访问的结点为当前结点的子节点时，当前结点出栈
			treeList = append(treeList, curr.Val)
			pre = curr
			stackList.Remove(e)
		}else{
			if curr.Right != nil {//先将右结点压栈
				stackList.PushBack(curr.Right)
			}
			if curr.Left != nil { //再将左结点入栈
				stackList.PushBack(curr.Left)
			}
		}
	}
	return treeList
}
