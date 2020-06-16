/*
   @Time : 2020/6/16 下午4:30
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("this reverse link list")

	//printList(generateList())

	head := generateList()
	printList(head)

	head = reverseList(head)
	printList(head)
}

type ListNode struct {
	Body int
	Next *ListNode
}

//双指针，遍历的同时改变指针即可。
//连续赋值好省事

func reverseList(head *ListNode) *ListNode {
	//if nil == head {
	//	return head
	//}
	//cur, next := head, head.Next
	//for ; nil != next; cur, next, next.Next = next, next.Next, cur {
	//}
	//head.Next = nil
	//return cur

	var next, newHead *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}

	return newHead
}

func printList(head *ListNode) {
	var index = head
	for index != nil {

		log.Println(index.Body)
		index = index.Next
	}
}

func generateList() (retList *ListNode) {
	var nums = []int{1, 2, 3, 4, 5}

	var head = new(ListNode)
	cur:= head

	//var index = head
	for i, v := range nums {
		node := new(ListNode)
		node.Body = v

		if i == 0 {
			head = node
			cur = node
		} else {
			cur.Next = node
			cur=node
		}
		//index.Next = node
		//index = node
	}
	return head
}
