package main

import "container/list"

type TreeNode struct {
   Val int
  Left *TreeNode
  Right *TreeNode
 }


/**
 *
 * @param root TreeNode类
 * @return int整型二维数组
 */
func levelOrder( root *TreeNode ) [][]int {
	// write code here
	r :=make([][]int,0)
	if root==nil{
		return r
	}
	l:=list.New()
	l.PushBack(root)
	for l.Len()>0{
		r1 :=make([]int,0)
		l1:=l.Len()
		for i:=0;i< l1 ;i++ {
			e:=l.Front()
			node:= e.Value.(*TreeNode)
			r1 =append(r1,node.Val )

			l.Remove(e)
			if node.Left!=nil{
				l.PushBack(node.Left)
			}
			if node.Right!=nil{
				l.PushBack(node.Right)
			}
		}
		r =append(r,r1)
	}
	return r
}

func levelOrder2( root *TreeNode ) [][]int {
	// write code here
	r :=make([][]int,0)
	if root==nil{
		return r
	}
	l1 :=make([]*TreeNode,0)
	l1 =append(l1,root)
	for len(l1)>0 {
		l2  :=make([]*TreeNode,0)
		r1 :=make([]int,0)
		for _,v:=range  l1 {
			r1 =append(r1,v.Val)
			if v.Left!=nil{
				l2=append(l2,v.Left)
			}
			if v.Right!=nil{
				l2=append(l2,v.Right)
			}
		}
		r=append(r,r1)
		l1=l2
	}
	return r
}