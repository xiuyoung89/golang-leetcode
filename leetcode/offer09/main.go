/*
   @Time : 2020/6/22 上午10:50
   @Author : young
   @File : main
   @Software: GoLand
*/
package main

import "log"

func main() {
	log.Println("offer 09 ")

	var tt = Constructor()
	log.Println(tt.DeleteHead())
	tt.AppendTail(10)
	tt.AppendTail(3)
	log.Println(tt.DeleteHead())
	log.Println(tt.DeleteHead())
	log.Println(tt.DeleteHead())
}

type stack []int

func (s *stack) Pop() int {
	n := len(*s)
	ret := (*s)[n-1]
	*s = (*s)[:n-1]
	return ret
}

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

type CQueue struct {
	//In  []int
	//Out []int
	In  stack
	Out stack
}

func Constructor() CQueue {
	return CQueue{}
	//newQueue := new(CQueue)
	//newQueue.In = make([]int, 0)
	//newQueue.Out = make([]int, 0)
	//return *newQueue
}

func (this *CQueue) AppendTail(value int) {
	//this.In = append(this.In, value)

	this.In.Push(value)
}

func (this *CQueue) DeleteHead() int {
	//
	//var ret = -1
	//
	//if len(this.Out) <= 0 {
	//	// 入‘出栈’
	//	for i := len(this.In) - 1; i >= 0; i-- {
	//		this.Out = append(this.Out, this.In[i])
	//	}
	//}
	//
	//if len(this.Out) <= 0 {
	//	return ret
	//}
	//
	//// delete
	//if len(this.In) > 1 {
	//	this.In = this.In[1:]
	//} else {
	//	this.In = []int{}
	//}
	//ret = this.Out[len(this.Out)-1]
	//this.Out = this.Out[0 : len(this.Out)-1]
	//return ret

	if len(this.Out) != 0 {
		return this.Out.Pop()
	} else if len(this.In) != 0 {
		for len(this.In) != 0 {
			this.Out.Push(this.In.Pop())
		}

		return this.Out.Pop()
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
