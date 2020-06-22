/*
   @Time : 2020/6/22 下午12:11
   @Author : young
   @File : main
   @Software: GoLand
*/
package main

import (
	"log"
	"math"
)

func main() {
	log.Println("offer 30")

	minStack := Constructor();
	//minStack.Push(0)
	//minStack.Push(1)
	//minStack.Push(0)
	//log.Println(minStack.Min())
	//minStack.Pop();
	//log.Println(minStack.Top())
	//log.Println(minStack.Min())
//["MinStack","push","push","push","top","pop","min","pop","min","pop","push","top","min","push","top","min","pop","min"]
//[[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]
//	minStack.Push(2147483646)
//	minStack.Push(2147483646)
//	minStack.Push(2147483647)

	minStack.Push(2)
	minStack.Push(0)
	minStack.Push(3)
	minStack.Push(0)
	log.Println(minStack.Min())
	minStack.Pop()
	log.Println(minStack.Min())
	minStack.Pop()
	log.Println(minStack.Min())
	minStack.Pop()
	log.Println(minStack.Min())



}

//var minInt = int(-(1 << 31))
//var maxInt = int(1<<31) - 1

type Stack []int

type MinStack struct {
	mainStack Stack
	helpStack Stack
}

func (s *Stack) Pop() int {
	if len(*s) != 0 {
		ret := (*s)[len(*s)-1]
		*s = (*s)[0 : len(*s)-1]
		return ret
	}
	return -1
}

func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

func (s *Stack) Top() int {
	if len(*s) != 0 {
		return (*s)[len(*s)-1]
	}else {
		return math.MaxInt32
	}
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		mainStack: []int{},
		helpStack: []int{},
	}
}

func (this *MinStack) Push(x int) {
	this.mainStack.Push(x)

	// help
	if x <= this.helpStack.Top() {
		this.helpStack.Push(x)
	}
}

func (this *MinStack) Pop() {
	x := this.mainStack.Top()
	this.mainStack.Pop()
	if x == this.helpStack.Top() {
		this.helpStack.Pop()
	}
}

func (this *MinStack) Top() int {
	return this.mainStack[len(this.mainStack)-1]
}

func (this *MinStack) Min() int {
	if this.mainStack.Top() < this.helpStack.Top() {
		return this.mainStack.Top()
	} else {
		return this.helpStack.Top()
	}
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
