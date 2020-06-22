/*
   @Time : 2020/6/22 下午2:55
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
	log.Println("offer 59")
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	log.Println(maxSlidingWindow(nums, k))
}

type helpQue []int

func (s *helpQue) AddTail(x int) {
	if len(*s) == 0 {
		*s = append(*s, x)
	} else {
		if x < (*s)[len(*s)-1] {
			*s = append(*s, x)
		}
	}
}

func (s *helpQue) PopLess(x int) {
	for len(*s) != 0 {
		if x > (*s)[len(*s)-1] {
			*s = (*s)[:len(*s)-1]
		} else {
			*s = append(*s, x)
			break
		}
	}

}

func (s *helpQue) PopHead(x int) {
	if len(*s) != 0 && x == (*s)[len(*s)-1] {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *helpQue) Head() int {
	if len(*s) != 0 {
		return (*s)[0]
	}
	return math.MinInt32
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := []int{}
	// 维护一张非严格单调队列
	helpQue := helpQue{}
	if len(nums) < k {
		return ret
	}
	// i and j
	for i, j := 0, 0; j < len(nums); {
		
		if j-i < k-1 {
			j++
		} else {
			ret = append(ret, helpQue.Head())
			i++
			j++
		}
	}
	return helpQue
}
