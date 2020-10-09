package main

func main() {
	head :=&ListNode{
		Val: 1,

	}

	reorderList (head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	li :=make([]*ListNode,0)
	next :=head
	for next!=nil {
		li = append(li,next)
		pre:=next
		next=next.Next
		pre.Next=nil
	}
	curr := head
	llen :=len(li)
	for i:=0;i<(llen+1)/2;i++ {
		if i!= 0{
			if curr==li[i]{
				break
			}
			curr.Next=li[i]
			curr=curr.Next
		}

		if curr==li[llen-i-1]{
			break
		}
		curr.Next= li[llen-i-1]
		curr=curr.Next
	}
}