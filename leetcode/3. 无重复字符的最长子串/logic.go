package main

import (
	"fmt"
)

func main() {
	res :=lengthOfLongestSubstring("abcabcbb")
	fmt.Println(res)
}
func lengthOfLongestSubstring(s string) int {
	var res int
	for i:=0;i<len(s);i++{
		rmap:=make(map[byte]struct{})
		for j:=i;j<len(s);j++{
			 if _,ok:=rmap[s[j]] ;ok{
			 	break
			 }
			rmap[s[j]]= struct{}{}
		}
		if len(rmap)>res {
			res=len(rmap)
		}
	}
	return res
}
