/*
   @Time : 2020/6/23 上午10:06
   @Author : young
   @File : main.go
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("offer 59-02 ")

	que := Constructor()

	log.Println(que.Pop_front())
	log.Println(que.Pop_front())
	log.Println(que.Pop_front())
	log.Println(que.Pop_front())
	log.Println(que.Pop_front())
	que.Push_back(15)
	log.Println(que.Max_value())
	que.Push_back(9)
	log.Println(que.Max_value())


	//que.Push_back(1)
	//que.Push_back(2)
	//log.Println(que.Max_value())
	//log.Println(que.Pop_front())
	//log.Println(que.Max_value())
}

type TQueue []int

func (q *TQueue) TopBack() int {
	if len(*q) != 0 {
		return (*q)[len(*q)-1]
	} else {
		return -1
	}
}

func (q *TQueue) PopBack() {
	if len(*q) != 0 {
		*q = (*q)[:len(*q)-1]
	}
}

func (q *TQueue) PushBack(v int) {
	*q = append(*q, v)
}

func (q *TQueue) TopHead() int {
	if len(*q) != 0 {
		return (*q)[0]
	} else {
		return -1
	}
}

func (q *TQueue) PopHead() int {
	pop := -1
	if len(*q) != 0 {
		pop = (*q)[0]
		*q = (*q)[1:]
	}
	return pop
}

type MaxQueue struct {
	MainQue TQueue
	HelpQue TQueue
}

func Constructor() MaxQueue {
	return MaxQueue{
		MainQue: []int{},
		HelpQue: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.HelpQue) != 0 {
		return this.HelpQue.TopHead()
	} else {
		return -1
	}
}

func (this *MaxQueue) Push_back(value int) {
	this.MainQue.PushBack(value)

	//检查辅助队列
	for len(this.HelpQue) != 0 && this.HelpQue.TopBack() < value {
		this.HelpQue.PopBack()
	}

	this.HelpQue.PushBack(value)

}

func (this *MaxQueue) Pop_front() int {
	pop := -1
	pop = this.MainQue.PopHead()

	if pop == this.HelpQue.TopHead() {
		this.HelpQue.PopHead()
	}
	return pop
}
