package main

import "fmt"

func main() {

	fmt.Println("333 Helloworld")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil{
		return head
	}

	nex := head
	back :=head
	for n>0 {
		n--
		back=back.Next
	}
	if back==nil {
		return head.Next
	}
	for back.Next!=nil {
		nex=nex.Next
		back=back.Next
	}

	nex.Next=nex.Next.Next
	return head
}
