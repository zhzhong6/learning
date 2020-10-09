package main

import "fmt"

func main() {

	fmt.Println("333 Helloworld")
}

type ListNode struct {
	Val  int
	Next *ListNode
}
func detectCycle(head *ListNode) *ListNode {
	slow :=head
	fast := head
	for {
		if fast ==nil || fast.Next==nil{
			return nil
		}
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast {
			break
		}
	}
 	slow=head
 	for {
		if slow==fast {
			return slow
		}
		slow=slow.Next
		fast=fast.Next

	}
}