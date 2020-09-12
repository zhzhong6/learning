package main


type ListNode struct{
  Val int
  Next *ListNode
}


/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */

func main()  {

}
func ReverseList( pHead *ListNode ) *ListNode {
	// write code here
	var r  *ListNode
	for pHead !=nil{
		n:=r
		r=pHead
		pHead=pHead.Next
		r.Next=n
	}
	return r
}