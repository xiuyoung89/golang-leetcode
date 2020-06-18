/*
@Time : 19-7-11 下午1:52
@Author : young
@File : addtwonumber
@Software: GoLand
*/
package main

import "log"

func main(){
	log.Println("add two number")
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	cur := res

	c := 0

	for l1 != nil || l2 != nil || c != 0 {

		if l1 != nil {
			cur.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			cur.Val += l2.Val
			l2 = l2.Next
		}

		cur.Val += c

		c = cur.Val/10
		cur.Val %= 10

		if l1 == nil && l2 == nil && c == 0 {
			return  res
		}

		cur.Next = &ListNode{}
		cur = cur.Next

	}

	return  nil
}
